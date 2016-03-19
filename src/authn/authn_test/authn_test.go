package authn_test

import (
	"authn"
	"util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authn", func() {
	var oldToken string
	var newToken string
	var err error
	var user = "Flowerchild"

	BeforeEach(func(){
		util.LoadConfigByName("test_config")
		authn.InitMint()
		authn.InitValidator()
	})

	Describe ("Token roundtrip", func(){
		Context("Freshly minted token", func(){
			It("Fresh token should be valid", func(){
				user, err := authn.Validate(authn.GetToken(user))
				Expect(err).To(BeNil())
				Expect(user).To(Equal(user))				
			})
		})		
	})

	Describe ("Reissuing a token", func(){

		Context("Username is preserved", func(){
			It("Should return the username that was given to the old token", func(){
				oldToken = authn.GetToken(user)
				newToken, err = authn.ReissueToken(oldToken)
				Expect(err).To(BeNil())

				u, err := authn.Validate(newToken)
				Expect(err).To(BeNil()) 		
				Expect(u).To(Equal(user))
			})
			
			It("An invalid token should not yield a username", func(){
				oldToken = "not a valid token at all"
				newToken, err = authn.ReissueToken(oldToken)
				Expect(err).To(Not(BeNil()))
			})
		})

	})
})
