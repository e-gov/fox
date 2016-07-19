#!/bin/bash

go test github.com/e-gov/fox/fox/fox_test \
	github.com/e-gov/fox/authn/authn_test \
	github.com/e-gov/fox/authz/authz_test \ 
	github.com/e-gov/fox/login/login_test
