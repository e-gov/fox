package login_test

import (
	. "authn"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	. "login"
	"net/http"
	"net/http/httptest"
	"net/url"
	"util"
	"authn"
)

var bufferLength int64 = 1048576

var _ = Describe("Login", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var username = "FantasticMrFox"
	var challenge = "test"
	var provider = "testprovider"

	BeforeEach(func() {
		util.LoadConfigByName("test_config")
		InitMint()
		InitValidator()
		router = NewRouter()
		recorder = httptest.NewRecorder()

		u, _ := url.Parse("/login")
		q := u.Query()
		q.Set("username", username)
		q.Set("challenge", challenge)
		q.Set("provider", provider)
		u.RawQuery = q.Encode()

		request, _ = http.NewRequest("GET", u.RequestURI(), nil)
	})

	Describe("Basic token generation", func() {
		Context("Login a user", func() {

			It("Should return 200", func() {
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))

				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())

				token := new(Token)

				err = json.Unmarshal(body, token)
				Expect(err).To(BeNil())

			})
		})
	})
	
	Describe("Token re-issue", func(){
		var token string
		BeforeEach(func(){
				authn.InitMint()
				token = authn.GetToken("testuser")

		})
		Context("Re-issue a token", func(){
			It("Should return 200", func(){
				request, _ = http.NewRequest("GET", "/reissue", nil)
				request.Header.Set("Authorization","Bearer " + token)
				
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})
			
			It("No token should yield 401", func(){
				request, _ = http.NewRequest("GET", "/reissue", nil)
				// No token is present
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(401))

			})
			It("False token should yield 401", func(){
				request, _ = http.NewRequest("GET", "/reissue", nil)
				request.Header.Set("Authorization","Bearer no such token")
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(401))
			})

		})

	})
})
