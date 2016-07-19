#!/bin/bash

mkdir -p coverage_in_html

go test github.com/e-gov/fox/fox/fox_test -coverpkg github.com/e-gov/fox/fox -coverprofile coverage_in_html/fox.out
go tool cover -html=coverage_in_html/fox.out -o coverage_in_html/fox_coverage.html

go test github.com/e-gov/fox/authz/authz_test -coverpkg github.com/e-gov/fox/authz -coverprofile coverage_in_html/authz.out
go tool cover -html=coverage_in_html/authz.out -o coverage_in_html/authz_coverage.html

go test github.com/e-gov/fox/authn/authn_test -coverpkg github.com/e-gov/fox/authn -coverprofile coverage_in_html/authn.out
go tool cover -html=coverage_in_html/authn.out -o coverage_in_html/authn_coverage.html

go test github.com/e-gov/fox/login/login_test -coverpkg github.com/e-gov/fox/login -coverprofile coverage_in_html/login.out
go tool cover -html=coverage_in_html/login.out -o coverage_in_html/login_coverage.html
