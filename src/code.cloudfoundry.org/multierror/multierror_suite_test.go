package multierror_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMultierror(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multierror Suite")
}
