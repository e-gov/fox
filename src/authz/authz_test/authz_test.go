package authz_test

import (
	"github.com/e-gov/fox/src/authz"

	"github.com/e-gov/fox/src/util"

	"github.com/e-gov/fox/src/fox"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("Authz", func() {

	var (
		provider authz.Provider
	)

	BeforeEach(func() {
		viper.Reset()
	})

	Describe("Getting the provider", func() {
		Context("Default provider is set", func() {
			It("Should return nil", func() {
				util.LoadConfigByPathWOExtension("authz/test_config_no-provider")
				provider = authz.GetProvider()
				Expect(provider).To(BeNil())
			})
		})

		Context("Default provider is set", func() {
			It("Should return simple provider", func() {
				util.LoadConfigByPathWOExtension("test_config")
				provider = authz.GetProvider()
				Expect(provider.GetName()).To(Equal("simple"))
			})
		})

	})

	Describe("Provider authorization queries", func() {
		BeforeEach(func() {
			util.LoadConfigByPathWOExtension("test_config")
			fox.NewRouter()
			provider = authz.GetProvider()
		})

		Context("User exists and has rights", func() {
			It("Should return true", func() {
				b := provider.IsAuthorized("fantasticmrfox", "GET", "/fox/foxes/{foxId}")
				Expect(b).To(Equal(true))
			})
		})
	})
})
