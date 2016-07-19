package authn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("authn_test")

func TestAuthn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authn suite")
}
