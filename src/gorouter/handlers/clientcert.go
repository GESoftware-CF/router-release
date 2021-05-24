package handlers

import (
	"encoding/pem"
	"net/http"
	"strings"

	"code.cloudfoundry.org/routing-release/gorouter/config"
	"code.cloudfoundry.org/routing-release/gorouter/errorwriter"
	"code.cloudfoundry.org/routing-release/gorouter/logger"
	"go.uber.org/zap"
	"github.com/urfave/negroni"
)

const xfcc = "X-Forwarded-Client-Cert"

type clientCert struct {
	skipSanitization  func(req *http.Request) bool
	forceDeleteHeader func(req *http.Request) (bool, error)
	forwardingMode    string
	logger            logger.Logger
	errorWriter       errorwriter.ErrorWriter
}

func NewClientCert(
	skipSanitization func(req *http.Request) bool,
	forceDeleteHeader func(req *http.Request) (bool, error),
	forwardingMode string,
	logger logger.Logger,
	ew errorwriter.ErrorWriter,
) negroni.Handler {
	return &clientCert{
		skipSanitization:  skipSanitization,
		forceDeleteHeader: forceDeleteHeader,
		forwardingMode:    forwardingMode,
		logger:            logger,
		errorWriter:       ew,
	}
}

func (c *clientCert) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	skip := c.skipSanitization(r)
	if !skip {
		switch c.forwardingMode {
		case config.FORWARD:
			if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
				r.Header.Del(xfcc)
			}
		case config.SANITIZE_SET:
			r.Header.Del(xfcc)
			if r.TLS != nil {
				replaceXFCCHeader(r)
			}
		}
	}

	delete, err := c.forceDeleteHeader(r)
	if err != nil {
		c.logger.Error("signature-validation-failed", zap.Error(err))
		c.errorWriter.WriteError(
			rw,
			http.StatusBadRequest,
			"Failed to validate Route Service Signature for x-forwarded-client-cert",
			c.logger,
		)
		return
	}
	if delete {
		r.Header.Del(xfcc)
	}
	next(rw, r)
}

func replaceXFCCHeader(r *http.Request) {
	// we only care about the first cert at this moment
	if len(r.TLS.PeerCertificates) > 0 {
		cert := r.TLS.PeerCertificates[0]
		b := pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
		certPEM := pem.EncodeToMemory(&b)
		r.Header.Add(xfcc, sanitize(certPEM))
	}
}

func sanitize(cert []byte) string {
	s := string(cert)
	r := strings.NewReplacer("-----BEGIN CERTIFICATE-----", "",
		"-----END CERTIFICATE-----", "",
		"\n", "")
	return r.Replace(s)
}
