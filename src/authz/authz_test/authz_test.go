package authz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"util"
	"authz"
	"fox"
)

var _ = Describe("Authz", func() {

	var provider authz.Provider

	BeforeEach(func(){
		util.LoadConfigByName("test_config")
		fox.NewRouter()
		provider = authz.GetProvider()
	})

	Describe ("Authorization querys", func(){
		BeforeEach(func(){
		})

		Context("User exists and has rights", func(){
			It("Should return true", func(){
				b := provider.IsAuthorized("fantasticmrfox", "GET", "/fox/foxes/{foxId}")
				Expect(b).To(Equal(true))
			})
		})
	})
})
