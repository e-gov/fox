package authz

// SimpleProvider is the struct to implement 
// the AuthZProvider interface on
type SimpleProvider struct{}

// IsAuthorized contains all the logic to authorize users to perform an action 
// on a URL. This includes, for example, mapping the method and URL to required
// roles on the remote LDAP server
func (provider SimpleProvider)IsAuthorized(user string, method string, url string) bool{
	log.Debugf("Authorized %s to access %s %s ", user, method, url)
	return true
}