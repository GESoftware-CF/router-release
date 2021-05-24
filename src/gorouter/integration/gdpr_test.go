package integration

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/routing-release/gorouter/route"
	"code.cloudfoundry.org/routing-release/gorouter/test"
	"code.cloudfoundry.org/routing-release/gorouter/test_util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

// Involves scrubbing client IPs, for more info on GDPR: https://www.eugdpr.org/
var _ = Describe("GDPR", func() {
	var testState *testState

	BeforeEach(func() {
		testState = NewTestState()
	})

	AfterEach(func() {
		if testState != nil {
			testState.StopAndCleanup()
		}
	})

	Context("when disable_log_forwarded_for is set true", func() {
		It("omits x-forwarded-for headers in log", func() {
			accessLog, err := ioutil.TempDir("", "accesslog")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(accessLog)

			testState.cfg.AccessLog.File = filepath.Join(accessLog, "access.log")

			testState.cfg.Logging.DisableLogForwardedFor = true
			testState.StartGorouterOrFail()

			testApp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer testApp.Close()

			hostname := "basic-app.some.domain"
			testState.register(testApp, hostname)

			req := testState.newRequest(fmt.Sprintf("http://%s", hostname))
			req.Header.Add("X-FORWARDED-FOR", "192.168.0.1")

			resp, err := testState.client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(testState.cfg.AccessLog.File).To(BeARegularFile())

			Eventually(func() ([]byte, error) {
				return ioutil.ReadFile(testState.cfg.AccessLog.File)
			}).Should(ContainSubstring(`x_forwarded_for:"-"`))

			f, err := ioutil.ReadFile(testState.cfg.AccessLog.File)
			Expect(err).NotTo(HaveOccurred())
			Expect(f).NotTo(ContainSubstring("192.168.0.1"))
		})

		It("omits x-forwarded-for from stdout", func() {
			testState.cfg.Status.Pass = "pass"
			testState.cfg.Status.User = "user"
			testState.cfg.Status.Port = 6705
			testState.cfg.Logging.DisableLogForwardedFor = true
			testState.StartGorouterOrFail()

			wsApp := test.NewWebSocketApp([]route.Uri{"ws-app." + test_util.LocalhostDNS}, testState.cfg.Port, testState.mbusClient, time.Millisecond, "")
			wsApp.Register()
			wsApp.Listen()
			wsApp.WaitUntilReady()

			routesURI := fmt.Sprintf("http://%s:%s@%s:%d/routes", testState.cfg.Status.User, testState.cfg.Status.Pass, "localhost", testState.cfg.Status.Port)

			Eventually(func() bool { return appRegistered(routesURI, wsApp) }, "2s").Should(BeTrue())

			conn, err := net.Dial("tcp", fmt.Sprintf("ws-app.%s:%d", test_util.LocalhostDNS, testState.cfg.Port))
			Expect(err).NotTo(HaveOccurred())

			x := test_util.NewHttpConn(conn)

			req := test_util.NewRequest("GET", "ws-app."+test_util.LocalhostDNS, "", nil)
			req.Header.Add("X-FORWARDED-FOR", "192.168.0.1")
			req.Header.Set("Upgrade", "websocket")
			req.Header.Set("Connection", "upgrade")
			x.WriteRequest(req)

			resp, _ := x.ReadResponse()
			Expect(resp.StatusCode).To(Equal(http.StatusSwitchingProtocols))

			x.WriteLine("hello from client")
			x.CheckLine("hello from server")

			x.Close()

			Eventually(gbytes.BufferReader(testState.gorouterSession.Out)).Should(gbytes.Say(`"X-Forwarded-For":"-"`))
			Expect(testState.gorouterSession.Out.Contents()).ToNot(ContainSubstring("192.168.0.1"))
		})
	})

	Context("when disable_log_source_ip is set true", func() {
		It("omits RemoteAddr in log", func() {
			accessLog, err := ioutil.TempDir("", "accesslog")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(accessLog)

			testState.cfg.AccessLog.File = filepath.Join(accessLog, "access.log")

			testState.cfg.Logging.DisableLogSourceIP = true
			testState.StartGorouterOrFail()

			testApp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer testApp.Close()

			hostname := "basic-app.some.domain"
			testState.register(testApp, hostname)

			req := testState.newRequest(fmt.Sprintf("http://%s", hostname))
			req.Header.Set("User-Agent", "foo-agent")

			resp, err := testState.client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(testState.cfg.AccessLog.File).To(BeARegularFile())

			Eventually(func() ([]byte, error) {
				return ioutil.ReadFile(testState.cfg.AccessLog.File)
			}).Should(ContainSubstring(`"foo-agent" "-"`))
		})

		It("omits RemoteAddr from stdout", func() {
			testState.cfg.Status.Pass = "pass"
			testState.cfg.Status.User = "user"
			testState.cfg.Status.Port = 6706
			testState.cfg.Logging.DisableLogSourceIP = true
			testState.StartGorouterOrFail()

			wsApp := test.NewWebSocketApp([]route.Uri{"ws-app." + test_util.LocalhostDNS}, testState.cfg.Port, testState.mbusClient, time.Millisecond, "")
			wsApp.Register()
			wsApp.Listen()
			wsApp.WaitUntilReady()

			routesURI := fmt.Sprintf("http://%s:%s@%s:%d/routes", testState.cfg.Status.User, testState.cfg.Status.Pass, "localhost", testState.cfg.Status.Port)

			Eventually(func() bool { return appRegistered(routesURI, wsApp) }, "2s").Should(BeTrue())

			conn, err := net.Dial("tcp", fmt.Sprintf("ws-app.%s:%d", test_util.LocalhostDNS, testState.cfg.Port))
			Expect(err).NotTo(HaveOccurred())

			x := test_util.NewHttpConn(conn)

			req := test_util.NewRequest("GET", "ws-app."+test_util.LocalhostDNS, "", nil)
			req.Header.Set("Upgrade", "websocket")
			req.Header.Set("Connection", "upgrade")
			x.WriteRequest(req)

			resp, _ := x.ReadResponse()
			Expect(resp.StatusCode).To(Equal(http.StatusSwitchingProtocols))

			x.WriteLine("hello from client")
			x.CheckLine("hello from server")

			x.Close()

			Eventually(gbytes.BufferReader(testState.gorouterSession.Out)).Should(gbytes.Say(`"RemoteAddr":"-"`))
		})
	})
})
