package authn_test

import (
	"authn"
	"util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authn", func() {
	var (
		oldToken  string
		newToken  string
		err       error
		user      = "FantasticMrFox"
		challenge = "test"
		provider  = "pwd"
	)

	BeforeEach(func() {
		util.LoadConfigByPathWOExtension("test_config")
		authn.InitMint()
		authn.InitValidator()
	})

	Describe("Token roundtrip", func() {
		Context("Freshly minted token", func() {
			It("Fresh token should be valid", func() {
				user, err := authn.Validate(authn.GetToken(user))
				Expect(err).To(BeNil())
				Expect(user).To(Equal(user))
			})
		})
		Context("Authenticating the user", func() {
			It("should return true, given valid username, challenge and provider", func() {
				booln := authn.Authenticate(user, challenge, provider)
				Expect(booln).To(BeTrue())
			})
		})
	})

	Describe("Reissuing a token", func() {
		Context("Username is preserved", func() {
			It("should return the username that was given to the old token", func() {
				oldToken = authn.GetToken(user)
				newToken, err = authn.ReissueToken(oldToken)
				Expect(err).To(BeNil())

				u, err := authn.Validate(newToken)
				Expect(err).To(BeNil())
				Expect(u).To(Equal(user))
			})

			It("An invalid token should not yield a username", func() {
				oldToken = "not a valid token at all"
				newToken, err = authn.ReissueToken(oldToken)
				Expect(err).To(Not(BeNil()))
			})
		})

	})
})
