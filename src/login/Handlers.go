package login

import (
	"net/http"
	"encoding/json"
	"authn"
	"log"
)

func sendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}


// Login returns a token 
func Login(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("username")
	c := r.FormValue("challenge")
	p := r.FormValue("provider")

	log.Println("u = " + u + ", c = " + c + ", p = " + p)
	
	// We need all three
	if u == "" || c == "" || p == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// If the credentials check out
	if authn.Authenticate(u, c, p){
		t := authn.GetToken(u)
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	
		if err := json.NewEncoder(w).Encode(authn.Token{Token: t}); err != nil {
			panic(err)
		}
	}else{
	// The credentials did not check out
		w.WriteHeader(http.StatusForbidden)		
	}

}
