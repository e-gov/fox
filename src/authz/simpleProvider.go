package authz

import(
	"strings"
	"authn"
)

// SimpleProvider is the struct to implement 
// the AuthZProvider interface on
type SimpleProvider struct{
	requirements []requirement
}

type requirement struct{
	role string
	method string
	url string
}

// IsAuthorized contains all the logic to authorize users to perform an action 
// on a URL. This includes, for example, mapping the method and URL to required
// roles on the remote LDAP server
// The current implementation allows any named user to access any restricted URL, 
// all unrestricted URLs are authorized to everyone. This is a naive implementation
// not to be used live
func (provider *SimpleProvider)IsAuthorized(user string, method string, url string) bool{

	for _, r := range provider.requirements{
		if strings.HasPrefix(url, r.url) && r.method == method{
			b := (user == "" && r.role == "*") || (user != "")
			log.Debugf("Request for %s to access %s %s returned %s", user, method, url, b)		
			return b
		}
	}
	log.Debugf("No matching rules found, denying %s to access %s %s ", user, method, url)
	return false
}

// AddRestriction adds restrictions to the current provider
func (provider *SimpleProvider)AddRestriction(role string, method string, url string){
	// Just take the prefix, remove the ID. Naive approach, do not replicate
	u := strings.Split(url, "{")[0]
	provider.requirements = append(provider.requirements, requirement{role, method, u})
	log.Debugf("Role %s mapped to %s  %s, %d rules in total", role, method, u, len(provider.requirements))
}

// GetRoles implements a naive role listing. All valid tokens will 
// result in a single "ADMIN" role, everybody else gets "*"
func (provider *SimpleProvider)GetRoles(token string)[]string{
	user, _ := authn.Validate(token)
	var roles []string
	
	if user != ""{
		return append(roles, "ADMIN")
	}
	
	return append(roles, "*") 
}