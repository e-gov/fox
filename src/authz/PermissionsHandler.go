package authz

import (
	"authn"
	"encoding/json"
	"net/http"
	"strings"
	"util"
)

// PermissionHandler validates the permissions of a user before further handling
func PermissionHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user string
		var ps string

		t := r.Header.Get("Authorization")
		if strings.HasPrefix(t, "Bearer ") {
			user, _ = authn.Validate(strings.SplitAfter(t, "Bearer ")[1])
			log.Debugf("Getting user %s from %s", user, t)
		} else {
			user = ""
			log.Debugf("Failed to get user from %s", t)
		}

		if GetProvider().IsAuthorized(user, r.Method, r.URL.RequestURI()) {
			log.Debugf("Authorized access, sending an error message")
			sw := util.MakeLogger(w)
			inner.ServeHTTP(sw, r)
		} else {
			log.Debugf("Unauthorized access, sending an error message")
			for _, p := range authn.KnownProviders() {
				if ps > "" {
					ps = ps + "," + p
				} else {
					ps = p
				}
			}
			w.Header().Set("WWW-Authenticate", "WWW-Authenticate:"+ps)
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(util.Error{Code: http.StatusUnauthorized, Message: "Permission denied"}); err != nil {
				panic(err)
			}
		}

	})
}
