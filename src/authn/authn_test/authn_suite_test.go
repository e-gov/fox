package authn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAuthn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authn Suite")
}
