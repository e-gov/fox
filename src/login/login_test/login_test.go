package login_test

import (
	. "login"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"io/ioutil"
	"io"
	"encoding/json"
)

var bufferLength int64 = 1048576

var _ = Describe("Login", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var keyName = "test_key.base64"
	var username = "testuser"
	
	BeforeEach(func(){
		LoadKeyByName(keyName)
		router = NewRouter()
		recorder = httptest.NewRecorder()
		
		request, _ = http.NewRequest("GET", "/login?username=" + username, nil)
		
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
