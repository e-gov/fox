package fox_test

import (
	"fox"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"log"
)

var _ = Describe("Fox", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	
	BeforeEach(func(){
		fox.LoadConfigByName("test_config.gcfg")
		router = fox.NewRouter("test")
		recorder = httptest.NewRecorder()
	})
	
	Describe ("GET /foxes", func(){
		BeforeEach(func(){
			request, _ = http.NewRequest("GET", "/fox/foxes", nil)
			log.Println("Request defined")
		})
		
		Context("Foxes exist", func(){
			It("Should return http 200", func(){
				log.Println("Serving http")
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))	
			})			
		})
	})


})
