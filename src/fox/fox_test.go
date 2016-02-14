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

var bufferLength int64 = 1048576

var _ = Describe("Fox", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var aFox Fox
	var anotherFox Fox
	
	BeforeEach(func(){
		LoadConfigByName("test_config.gcfg")
		router = NewRouter("test")
		recorder = httptest.NewRecorder()
		
		aFox = Fox{
			Name: "Rein",
			Parents: []string{"2", "3"},			
		}
		
		anotherFox = Fox{
			Name: "NewName",
			Parents: []string{"4", "5"},
		}
	})
	
	Describe("Adding foxes", func(){
		BeforeEach(func(){
			m, _ := json.Marshal(aFox)
			request, _ = http.NewRequest("POST", "/fox/foxes", bytes.NewReader(m))
		})
		
		Context("Adding, reading and updating a fox", func(){
			// Simple adding of a fox
			It("Should return 201", func(){
				var f Fox
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(201))				

			// See if we can get the same fox back
			// Read the UUID from the response first
				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())

				var id *UUID
				id = new(UUID) 
				
				err = json.Unmarshal(body, id)
				Expect(err).To(BeNil())

				f = getFox(id.Uuid, router)				
				
				// Updating the fox we just received
				anotherFox.Uuid = id.Uuid
				m, _ := json.Marshal(anotherFox)
				request, _ = http.NewRequest("PUT", "/fox/foxes/" + id.Uuid, bytes.NewReader(m))
				
				recorder = httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(202))
				
				// Read the fox again
				f = getFox(id.Uuid, router)
				Expect(Compare(f, anotherFox)).To(BeTrue())
			})
		})
	})
	
	Describe ("Reading the fox list", func(){
		BeforeEach(func(){
			request, _ = http.NewRequest("GET", "/fox/foxes", nil)
		})
		
		Context("Foxes exist", func(){
			It("Should return http 200", func(){
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))	
			})
		})
		
		Context("Random fox should return 404", func(){
			It("Should return 404", func(){
				request, _ = http.NewRequest("GET", "/fox/foxes/nosuchfoxforsure", nil)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(404))
			})
		})
	})

	Describe ("Updating a fox", func(){
		var m []byte
		
		BeforeEach(func(){
				m, _ = json.Marshal(aFox)			
		})
		
		Context("Update a fox", func(){
			It("Should return 201", func(){				
				request, _ = http.NewRequest("PUT", "/fox/foxes/nosuchfoxforsure", bytes.NewReader(m))
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(201))
			})
			
			It("Should return 422", func(){
				request, _ = http.NewRequest("POST", "/fox/foxes", bytes.NewReader(m))	
				// Read the UUID from the response first
				router.ServeHTTP(recorder, request)
				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())

				var id *UUID
				id = new(UUID) 
				
				err = json.Unmarshal(body, id)
				Expect(err).To(BeNil())

				recorder = httptest.NewRecorder()
				m, _ = json.Marshal("This is not a valid Fox")
				request, _ = http.NewRequest("PUT", "/fox/foxes/" + id.Uuid, bytes.NewReader(m))

				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(422))
			})

		})
	})

})

func getFox(uuid string, router *mux.Router) Fox{
	var r *http.Request
	var f *Fox
	
	recorder := httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/fox/foxes/" + uuid, nil)
	router.ServeHTTP(recorder, r)
	Expect(recorder.Code).To(Equal(200))

	body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
	Expect(err).To(BeNil())
				
				
	f = new(Fox)
				
	err = json.Unmarshal(body, f)
	Expect(err).To(BeNil())

	return *f
}