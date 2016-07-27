#=============== FoxService ===============
FROM golang:1.6

RUN mkdir -p /src/github.com/e-gov/fox/fox/

COPY ./src/github.com/e-gov/fox/fox/ /src/github.com/e-gov/fox/fox/
COPY ./config /config/

RUN mkdir -p /src/vendor/github.com/e-gov/fox/

WORKDIR /src/vendor/github.com/e-gov/fox/
 
COPY ./src/github.com/e-gov/fox/authn/ authn/
COPY ./src/github.com/e-gov/fox/authz/ authz/
COPY ./src/github.com/e-gov/fox/util/ util/

ENV GOPATH=/

WORKDIR /src/github.com/e-gov/fox/fox/foxservice/

# fox,authn,authz use this config 
# temporary database data for foxservice
RUN mkdir -p /src/github.com/e-gov/fox/tmp/foxdb/ \
	&& go get -d && CGO_ENABLED=0 go build -a -installsuffix cgo -v \
	-o foxservice.o FoxService.go \
	&& cp foxservice.o /bin/foxservice

CMD ["/bin/foxservice"]

