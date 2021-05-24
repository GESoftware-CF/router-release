package testrunner

import (
	"os/exec"
	"strconv"
	"time"

	"github.com/tedsuo/ifrit/ginkgomon"
)

// Args used by runner
type Args struct {
	BaseLoadBalancerConfigFilePath string
	LoadBalancerConfigFilePath     string
	ConfigFilePath                 string
	RoutingGroupCheckExit          bool
}

func (args Args) ArgSlice() []string {
	return []string{
		"-tcpLoadBalancerConfig=" + args.LoadBalancerConfigFilePath,
		"-tcpLoadBalancerBaseConfig=" + args.BaseLoadBalancerConfigFilePath,
		"-haproxyReloader=fixtures/fake_haproxy_reloader",
		"-config=" + args.ConfigFilePath,
		"-tokenFetchRetryInterval", "1s",
		"-staleRouteCheckInterval", "5s",
		"-logLevel=debug",
		"-routingGroupCheckExit=" + strconv.FormatBool(args.RoutingGroupCheckExit),
	}
}

func New(binPath string, args Args) *ginkgomon.Runner {
	return ginkgomon.New(ginkgomon.Config{
		Name:              "tcp-router",
		AnsiColorCode:     "1;97m",
		StartCheck:        "tcp-router.started",
		StartCheckTimeout: 10 * time.Second,
		Command:           exec.Command(binPath, args.ArgSlice()...),
	})
}
