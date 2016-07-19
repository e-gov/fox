package login_test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/e-gov/fox/login"
	"github.com/e-gov/fox/authn"

	"github.com/e-gov/fox/util"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var bufferLength int64 = 1048576

var _ = Describe("Login", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var username = "FantasticMrFox"
	var challenge = "test"
	var provider = "testprovider"

	var token string

	BeforeEach(func() {
		util.LoadConfigByPathWOExtension("test_config")
		authn.InitMint()
		authn.InitValidator()
		router = login.NewRouter()
		recorder = httptest.NewRecorder()

		u, _ := url.Parse("/login")
		q := u.Query()
		q.Set("username", username)
		q.Set("challenge", challenge)
		q.Set("provider", provider)
		u.RawQuery = q.Encode()

		request, _ = http.NewRequest("GET", u.RequestURI(), nil)

		authn.InitMint()
		token = authn.GetToken("testuser")

	})

	Describe("Basic token generation", func() {
		Context("Login a user", func() {

			It("Should return 200", func() {
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))

				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())

				token := new(authn.Token)

				err = json.Unmarshal(body, token)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Token re-issue", func() {
		Context("Re-issue a token", func() {
			It("Should return 200", func() {
				request, _ = http.NewRequest("GET", "/login/reissue", nil)
				request.Header.Set("Authorization", "Bearer "+token)

				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("No token should yield 401", func() {
				request, _ = http.NewRequest("GET", "/login/reissue", nil)
				// No token is present
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(401))

			})
			It("False token should yield 401", func() {
				request, _ = http.NewRequest("GET", "/login/reissue", nil)
				request.Header.Set("Authorization", "Bearer no such token")
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(401))
			})

		})

	})

	Describe("Role list", func() {
		Context("Get a role list without a token", func() {
			It("Should get a brief list of roles", func() {
				request, _ = http.NewRequest("GET", "/login/roles", nil)

				Expect(getRoles(router, recorder, request)).To(Not(BeEmpty()))
			})
			It("Should get an admin role", func() {
				request, _ = http.NewRequest("GET", "/login/roles", nil)
				request.Header.Set("Authorization", "Bearer "+token)

				Expect(getRoles(router, recorder, request)).To(ContainElement("ADMIN"))
			})
		})
	})

	Describe("Getting statistics", func() {
		Context("Get the stats", func() {
			It("Should return 200", func() {
				request, _ := http.NewRequest("GET", "/login/status", nil)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})
		})
	})

})

func getRoles(router *mux.Router, recorder *httptest.ResponseRecorder, request *http.Request) []string {
	var roles []string
	router.ServeHTTP(recorder, request)
	Expect(recorder.Code).To(Equal(200))

	body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
	Expect(err).To(BeNil())

	err = json.Unmarshal(body, &roles)
	Expect(err).To(BeNil())

	return roles
}
