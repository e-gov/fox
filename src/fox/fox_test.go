package fox_test

import (
	. "fox"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"io"
)

var _ = Describe("Fox", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var aFox Fox
	
	BeforeEach(func(){
		LoadConfigByName("test_config.gcfg")
		router = NewRouter("test")
		recorder = httptest.NewRecorder()
		
		aFox = Fox{
			Name: "Rein",
			Parents: []string{"2", "3"},			
		}
	})
	
	Describe("Adding foxes", func(){
		BeforeEach(func(){
			m, err := json.Marshal(aFox)
			Expect(err).To(BeNil())
			request, _ = http.NewRequest("POST", "/fox/foxes", bytes.NewReader(m))
		})
		
		Context("Adding a fox", func(){
			// Simple adding of a fox
			It("Should return 201", func(){
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(201))				

			// See if we can get the same fox back
				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, 1048576))
				Expect(err).To(BeNil())

//				fmt.Printf(">>> got body %d", len(recorder.Body.Bytes()))
				var id *UUID
				id = new(UUID) 
				
				err = json.Unmarshal(body, id)
				Expect(err).To(BeNil())
				request, _ = http.NewRequest("GET", "/fox/foxes/" + id.Uuid, nil)
				
				recorder = httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))					
			})
							
			
			// Send garbage instead of a Fox
			It("Should return 422", func(){
				m, err := json.Marshal("This is not a valid Fox")
				Expect(err).To(BeNil())
				request, _ = http.NewRequest("POST", "/fox/foxes", bytes.NewReader(m))

				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(422))
			})
		})
	})
	
	Describe ("GET /foxes", func(){
		BeforeEach(func(){
			request, _ = http.NewRequest("GET", "/fox/foxes", nil)
		})
		
		Context("Foxes exist", func(){
			It("Should return http 200", func(){
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))	
			})			
		})
	})


})
