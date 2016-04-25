package login

import (
	"net/http"
	"encoding/json"
	"authn"
	"strings"
	"authz"
)

func sendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}



// Login returns a token 
func Login(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("username")
	c := r.FormValue("challenge")
	p := r.FormValue("provider")

	// We need all three
	if u == "" || c == "" || p == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If the credentials check out
	if authn.Authenticate(u, c, p) {
		sendToken(w, u)
	} else {
		// The credentials did not check out
		w.WriteHeader(http.StatusForbidden)
	}

}

// Reissue re-issues a new token based on an existing valid one
func Reissue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	t := r.Header.Get("Authorization")
	if strings.HasPrefix(t, "Bearer ") {
		user, err := authn.Validate(strings.SplitAfter(t, "Bearer ")[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			sendToken(w, user)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}

// Roles returns a list of applicable roles based on the username in the token
func Roles(w http.ResponseWriter, r *http.Request) {
	var token string

	t := r.Header.Get("Authorization")
	if strings.HasPrefix(t, "Bearer ") {
		token = strings.SplitAfter(t, "Bearer ")[1]
	} else {
		token = ""
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(authz.GetProvider().GetRoles(token)); err != nil {
		panic(err)
	}

}

// Mint a token for a given username and send it to the specified writer
func sendToken(w http.ResponseWriter, u string) {
	t := authn.GetToken(u)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(authn.Token{Token: t}); err != nil {
		panic(err)
	}

}