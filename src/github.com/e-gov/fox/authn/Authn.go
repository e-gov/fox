package authn

import (
	"errors"
)

var (
	confVersion int
)

type TokenStruct struct {
	Username string `json:"username"`
	MintTime string `json:"mintTime"`
}

type Token struct {
	Token string `json:"token"`
}

// Authenticate sends the username and challenge to the authentication
// provider requested and passes the resulting boolean back.
// Any non-true result including technical failures and the authentication
// provider being unknown will be interepreted
// as the user not being authenticated
func Authenticate(username string, challenge string, provider string) bool {
	return GetProvider(provider).Authenticate(username, challenge)
}

// Validate the Fernet token and extract the username
// Returns either a username or an error
func Validate(token string) (string, error) {

	t := Decrypt(token)
	if t == nil {
		return "", errors.New("Invalid token")
	}
	return t.Username, nil
}

// KnownProviders returns the list of known authentication providers
func KnownProviders() []string {
	var s []string
	return append(s, "password")
}
