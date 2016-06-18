#!/bin/bash

mkdir -p coverage_in_html

go test fox/fox_test -coverpkg fox -coverprofile coverage_in_html/fox.out
go tool cover -html=coverage_in_html/fox.out -o coverage_in_html/fox_coverage.html

go test authz/authz_test -coverpkg authz -coverprofile coverage_in_html/authz.out
go tool cover -html=coverage_in_html/authz.out -o coverage_in_html/authz_coverage.html

go test authn/authn_test -coverpkg authn -coverprofile coverage_in_html/authn.out
go tool cover -html=coverage_in_html/authn.out -o coverage_in_html/authn_coverage.html

go test login/login_test -coverpkg login -coverprofile coverage_in_html/login.out
go tool cover -html=coverage_in_html/login.out -o coverage_in_html/login_coverage.html
