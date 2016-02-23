package login

import (
	"net/http"
//	fernet "github.com/fernet/fernet-go"
//	"log"
	"encoding/json"
)

func sendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}


// Login returns a token 
func Login(w http.ResponseWriter, r *http.Request) {
	
	s := GetToken(r.FormValue("username"))
	if err := json.NewEncoder(w).Encode(Token{Token: s}); err != nil {
		panic(err)
	}

}
