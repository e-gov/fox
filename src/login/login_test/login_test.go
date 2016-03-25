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

	var token string

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

				token := new(Token)

				err = json.Unmarshal(body, token)
				Expect(err).To(BeNil())
			})
		})
	})
	
	Describe("Token re-issue", func(){
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
	
	Describe("Role list", func(){
		Context("Get a role list without a token", func(){
			It("Should get a brief list of roles", func(){	
				request, _ = http.NewRequest("GET", "/roles", nil)

				Expect(getRoles(router, recorder, request)).To(Not(BeEmpty()))
			})
			It("Should get an admin role", func(){	
				request, _ = http.NewRequest("GET", "/roles", nil)
				request.Header.Set("Authorization","Bearer " + token)
				
				Expect(getRoles(router, recorder, request)).To(ContainElement("ADMIN"))
			})
		})
	})

})

func getRoles(router *mux.Router, recorder *httptest.ResponseRecorder, request *http.Request)[]string{
	var roles []string			
	router.ServeHTTP(recorder, request)
	Expect(recorder.Code).To(Equal(200))
	
	body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
	Expect(err).To(BeNil())

	err = json.Unmarshal(body, &roles)
	Expect(err).To(BeNil())

	return roles
}

