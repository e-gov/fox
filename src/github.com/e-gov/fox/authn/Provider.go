package authn

import (
	pwd "github.com/e-gov/fox/authn/pwd"
)

// Provider is a simple interface all authentication providers must implement
type Provider interface {
	Authenticate(string, string) bool
}

func GetProvider(name string) Provider {
	return new(pwd.PwdProvider)
}
