package authz

import (
	"crypto/tls"
	"fmt"
	"path"
	"strings"

	"github.com/e-gov/fox/authn"
	"github.com/e-gov/fox/util"

	"gopkg.in/ldap.v2"

	log "github.com/Sirupsen/logrus"
)

type LdapProvider struct {
	requirements []requirement
}

// TODO: multiple groups for users
// TODO: make use of the registryGroups organizational unit in the LDAP directory
func (provider *LdapProvider) IsAuthorized(user string, method string, url string) bool {
	entry, err := provider.getEntryForUser(user)

	if err != nil {
		log.Fatal(err)
		return false
	}

	for _, r := range provider.requirements {
		if strings.HasPrefix(url, r.url) && r.method == method && r.role == entry.GetAttributeValue("cn") {
			b := (user == "" && r.role == "registry user") || (user != "")
			log.Debugf("Request for %s to access %s %s returned %t", user, method, url, b)
			return b
		}
	}

	log.Debugf("No matching rules found, denying %s to access %s %s ", user, method, url)
	return false
}

func (provider *LdapProvider) AddRestriction(role string, method string, url string) {
	u := path.Dir(url)
	provider.requirements = append(provider.requirements, requirement{role, method, u})
	log.Debugf("Role %s mapped to %s  %s, %d rules in total", role, method, u, len(provider.requirements))
}

// TODO: make it possible to get multiple groups for users
func (provider *LdapProvider) GetRoles(token string) []string {
	user, _ := authn.Validate(token)

	entry, err := provider.getEntryForUser(user)

	if err != nil {
		return []string{"registry user"}
	} else {
		return []string{entry.GetAttributeValue("cn")}
	}

}

func (provider *LdapProvider) getEntryForUser(user string) (ldap.Entry, error) {
	var entry ldap.Entry
	ldapConnection, err := provider.startTlsConnectionToLdapServer()

	if err != nil {
		return entry, err
	}

	defer ldapConnection.Close()

	_, err = provider.bindFoxApiAsUser(ldapConnection)
	if err != nil {
		return entry, err
	}

	entry, err = provider.searchForRegistryUserEntry(ldapConnection, user)
	if err != nil {
		return entry, err
	}
	return entry, nil

}

func (provider *LdapProvider) startTlsConnectionToLdapServer() (*ldap.Conn, error) {
	ldapConnection, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 10389))
	if err != nil {
		return nil, err
	}

	// Reconnect with TLS.
	// InsecureSkipVerify controls whether a client verifies the
	// server's certificate chain and host name.
	// If InsecureSkipVerify is true, TLS accepts any certificate
	// presented by the server and any host name in that certificate.
	// In this mode, TLS is susceptible to man-in-the-middle attacks.
	// This should be used only for testing.
	err = ldapConnection.StartTLS(&tls.Config{InsecureSkipVerify: true})

	if err != nil {
		return nil, err
	} else {
		return ldapConnection, nil
	}
}

func (provider *LdapProvider) bindFoxApiAsUser(ldapConnection *ldap.Conn) (bool, error) {
	err := ldapConnection.Bind(
		fmt.Sprintf("uid=%s,ou=users,ou=system", util.GetConfig().Authz.LDAPProvider.User),
		util.GetConfig().Authz.LDAPProvider.Password)
	fmt.Print("")
	if err != nil {
		return true, err
	} else {
		return false, nil
	}
}

func (provider *LdapProvider) searchForRegistryUserEntry(ldapConnection *ldap.Conn, user string) (ldap.Entry, error) {
	var userEntry ldap.Entry
	searchRequest := ldap.NewSearchRequest(
		"ou=registryUsers,o=e-gov",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=person))",
		[]string{"uid", "cn"},
		nil, // controls -- eg, password policies.
	)

	sr, err := ldapConnection.Search(searchRequest)
	if err != nil {
		return userEntry, err
	}

	for _, entry := range sr.Entries {
		if entry.GetAttributeValue("uid") == user {
			userEntry = *entry
		}
	}

	if err != nil {
		return userEntry, err
	}

	return userEntry, nil
}

// GetName returns the name of the provider, ldap in this case
func (provider *LdapProvider) GetName() string {
	return "ldap"
}
