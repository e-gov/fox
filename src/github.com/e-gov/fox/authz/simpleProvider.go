package authz

import (
	"strings"

	"github.com/e-gov/fox/authn"

	log "github.com/Sirupsen/logrus"
)

// SimpleProvider is the struct to implement
// the AuthZProvider interface on
type SimpleProvider struct {
	requirements []requirement
}

type requirement struct {
	role   string
	method string
	url    string
}

// IsAuthorized contains all the logic to authorize users to perform an action
// on a URL. This includes, for example, mapping the method and URL to required
// roles on the remote LDAP server
// The current implementation allows any named user to access any restricted URL,
// all unrestricted URLs are authorized to everyone. This is a naive implementation
// not to be used live
func (provider *SimpleProvider) IsAuthorized(user string, method string, url string) bool {

	for _, r := range provider.requirements {
		log.WithFields(log.Fields{
			"path": url,
			"method": method,
			"user": user,
		}).Debugf("Comparing user %s, %s:%s to %s:%s", user, method, url, r.method, r.url)
		if strings.HasPrefix(url, r.url) && r.method == method {
			b := (user == "" && r.role == "*") || (user != "")
			log.WithFields(log.Fields{
				"path": url,
				"method": method,
				"user": user,
			}).Debugf("Request for user %s to access %s %s returned %t", user, method, url, b)
			return b
		}
	}
	// Unless there is an explicit rule about allowing access, we fall back to denying access
	log.WithFields(log.Fields{
		"path": url,
		"method": method,
		"user": user,
	}).Infof("No matching rules found, denying %s to access %s %s ", user, method, url)
	return false
}

// AddRestriction adds restrictions to the current provider
func (provider *SimpleProvider) AddRestriction(role string, method string, url string) {
	// Just take the prefix, remove the ID. Naive approach, do not replicate
	u := strings.Split(url, "{")[0]
	provider.requirements = append(provider.requirements, requirement{role, method, u})
	log.WithFields(log.Fields{
		"path": u,
		"method": method,
		"role": role,
	}).Debugf("Role %s mapped to %s  %s, %d rules in total", role, method, u, len(provider.requirements))
}

// GetRoles implements a naive role listing for a user. All valid tokens will
// result in a single "ADMIN" role, everybody else gets "*"
func (provider *SimpleProvider) GetRoles(token string) []string {
	user, _ := authn.Validate(token)
	var roles []string

	if user != "" {
		return append(roles, "ADMIN")
	}

	return append(roles, "*")
}

// GetName returns the name of the provider
func (provider *SimpleProvider) GetName() string {
	return "simple"
}
