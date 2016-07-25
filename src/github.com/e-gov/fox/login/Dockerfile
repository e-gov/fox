#=============== LoginService ===============
FROM golang:1.6

RUN mkdir -p /src/github.com/e-gov/fox/login/

COPY ./src/github.com/e-gov/fox/login/ /src/github.com/e-gov/fox/login/
COPY ./config /config/

RUN mkdir -p /src/vendor/github.com/e-gov/fox/

WORKDIR /src/vendor/github.com/e-gov/fox/
 
COPY ./src/github.com/e-gov/fox/authn/ authn/
COPY ./src/github.com/e-gov/fox/authz/ authz/
COPY ./src/github.com/e-gov/fox/util/ util/

ENV GOPATH=/

# authn,authz use this config
# getConfigFolderPath() should find the config in /config/ 
RUN mkdir /config/docker \
	&& cp -n /config/config.json.template /config/docker/config.json \
	&& go get -d vendor/github.com/e-gov/fox/authn/keygen/ \ 
	&& go run /src/vendor/github.com/e-gov/fox/authn/keygen/KeyGen.go > /config/key.base64 

WORKDIR /src/github.com/e-gov/fox/login/loginservice/

RUN go get -d && CGO_ENABLED=0 go build -a -installsuffix cgo -v \
	-o loginservice.o LoginService.go \
	&& cp loginservice.o /bin/loginservice

CMD ["/bin/loginservice"]

