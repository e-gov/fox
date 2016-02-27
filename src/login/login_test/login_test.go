package login_test

import (
	. "login"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "authn"
	"net/http"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"io/ioutil"
	"io"
	"encoding/json"
	"net/url"
)

var bufferLength int64 = 1048576

var _ = Describe("Login", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var keyName = "test_key.base64"
	var username = "testuser"
	var challenge = "This might be a password or some token"
	var provider = "testprovider"
	
	BeforeEach(func(){
		LoadKeyByName(keyName)
		router = NewRouter()
		recorder = httptest.NewRecorder()
		
//		Error{Code: 404, Message: "Fox list not available"}
//		m, _ := json.Marshal(anotherFox)
//		request, _ = http.NewRequest("PUT", "/fox/foxes/" + foxID, bytes.NewReader(m))				

		u, _ := url.Parse("/login")
		q := u.Query()
		q.Set("username", username)
		q.Set("challenge", challenge)
		q.Set("provider", provider)
		u.RawQuery = q.Encode()		

		request, _ = http.NewRequest("GET", u.RequestURI(), nil)
		
	})
	
	Describe("Basic token generation", func(){
		Context("Login a user", func(){

			It("Should return 200", func(){
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
})
