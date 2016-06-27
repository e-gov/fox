package fox_test

import (
	"authn"
	"bytes"
	"encoding/json"
	. "fox"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"util"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//	"fmt"
	"time"
)

var bufferLength int64 = 1048576
var token string

var _ = Describe("Fox", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var aFox Fox
	var anotherFox Fox

	BeforeEach(func() {
		util.LoadConfigByPathWOExtension("test_config")
		router = NewRouter()
		recorder = httptest.NewRecorder()

		aFox = Fox{
			Name:    "Rein",
			Parents: []string{"2", "3"},
		}

		anotherFox = Fox{
			Name:    "NewName",
			Parents: []string{"4", "5"},
		}
		authn.InitValidator()
		authn.InitMint()
		token = authn.GetToken("testuser")
	})

	Describe("Adding foxes", func() {
		Context("Adding, reading and updating a fox", func() {
			// Simple adding of a fox
			It("Should return 201", func() {

				// Add the fox
				foxID := addFox(aFox, router)

				// See if we can find it
				_ = getFox(foxID, router)
			})
		})
	})

	Describe("Reading the fox list", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/fox/foxes", nil)
		})

		Context("Foxes exist", func() {
			It("Should return http 200", func() {
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})
		})

		Context("Random fox should return 404", func() {
			It("Should return 404", func() {
				request, _ = http.NewRequest("GET", "/fox/foxes/nosuchfoxforsure", nil)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(404))
			})
		})
	})

	Describe("Updating a fox", func() {
		var m []byte

		BeforeEach(func() {
			m, _ = json.Marshal(aFox)
			recorder = httptest.NewRecorder()
		})

		Context("Update a fox", func() {
			It("Should return 202", func() {
				// Add the fox
				foxID := addFox(aFox, router)

				// Attempt to update it
				anotherFox.Uuid = foxID
				m, _ := json.Marshal(anotherFox)
				request, _ = http.NewRequest("PUT", "/fox/foxes/"+foxID, bytes.NewReader(m))
				request.Header.Set("Authorization", "Bearer "+token)

				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(202))

				// Read the fox back
				f := getFox(foxID, router)
				Expect(Compare(f, anotherFox)).To(BeTrue())
			})

			It("Should return 201", func() {
				request, _ = http.NewRequest("PUT", "/fox/foxes/nosuchfoxforsure", bytes.NewReader(m))
				request.Header.Set("Authorization", "Bearer "+token)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(201))
			})

			It("Should return 422", func() {
				var s string
				s = addFox(aFox, router)

				m, _ = json.Marshal("This is not a valid Fox")
				request, _ = http.NewRequest("PUT", "/fox/foxes/"+s, bytes.NewReader(m))
				request.Header.Set("Authorization", "Bearer "+token)

				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(422))
			})

		})
	})

	Describe("Deleting a fox", func() {
		BeforeEach(func() {
			recorder = httptest.NewRecorder()
		})

		Context("Delete a fox", func() {
			It("Should return 200", func() {
				foxID := addFox(aFox, router)
				request, _ := http.NewRequest("DELETE", "/fox/foxes/"+foxID, nil)
				request.Header.Set("Authorization", "Bearer "+token)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("Should return 404", func() {
				request, _ := http.NewRequest("DELETE", "/fox/foxes/nosuchfoxforsure", nil)
				request.Header.Set("Authorization", "Bearer "+token)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(404))
			})
		})
	})

	Describe("Permissions should be checked", func() {
		It("Should return 401", func() {
			// No token present
			request, _ := http.NewRequest("DELETE", "/fox/foxes/nosuchfoxforsure", nil)
			router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(401))

			Expect(recorder.HeaderMap).To(ContainElement(ContainElement(ContainSubstring("WWW-Authenticate"))))
		})

		It("Should fail if we wait until the token expires", func(done Done) {
			time.Sleep(5 * time.Second)
			addFoxExpectingCode(aFox, router, 401)
			close(done)
		}, 180)
	})

	Describe("Getting statistics", func() {
		Context("Get the stats", func() {
			It("Should return 200", func() {
				request, _ := http.NewRequest("GET", "/fox/status", nil)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})
		})
	})
})

func addFox(f Fox, router *mux.Router) string {
	return addFoxExpectingCode(f, router, 201)
}

func addFoxExpectingCode(f Fox, router *mux.Router, retCode int) string {
	var id *UUID
	var r *http.Request

	m, _ := json.Marshal(f)

	recorder := httptest.NewRecorder()
	r, _ = http.NewRequest("POST", "/fox/foxes", bytes.NewReader(m))
	r.Header.Set("Authorization", "Bearer "+token)

	router.ServeHTTP(recorder, r)
	Expect(recorder.Code).To(Equal(retCode))

	body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
	Expect(err).To(BeNil())

	id = new(UUID)

	err = json.Unmarshal(body, id)
	Expect(err).To(BeNil())

	return id.Uuid
}

func getFox(uuid string, router *mux.Router) Fox {
	var r *http.Request
	var f *Fox

	recorder := httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "/fox/foxes/"+uuid, nil)
	router.ServeHTTP(recorder, r)
	Expect(recorder.Code).To(Equal(200))

	body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
	Expect(err).To(BeNil())

	f = new(Fox)

	err = json.Unmarshal(body, f)
	Expect(err).To(BeNil())

	return *f
}
