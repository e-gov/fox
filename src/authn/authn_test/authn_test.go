package authn_test

import (
	"authn"
	"fox"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authn", func() {
	var oldToken string
	var newToken string

	Describe ("Reissuing a token", func(){
		BeforeEach(func(){
			fox.LoadConfig()
			authn.InitMint()
			authn.InitValidator()
		})

		Context("Username is preserved", func(){
			It("Should return the username that was given to the old token", func(){
				oldToken = authn.GetToken("Flowerchild")
				newToken = authn.ReissueToken(oldToken)
				Expect(authn.Validate(newToken)).To(Equal("Flowerchild"))
			})
		})

	})
})
