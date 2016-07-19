#=============== FoxService ===============
FROM golang:1.6

COPY ./src/fox/ /fox/

ENV GOPATH=/fox/vendor

WORKDIR /fox/foxservice/

RUN mkdir /fox/vendor && go get \ 
	&& CGO_ENABLED=0 go build -a -installsuffix cgo -v \
	-o foxservice.o FoxService.go

COPY foxservice.o /bin/foxservice

CMD ["bin/foxservice"]

