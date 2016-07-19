#=============== LoginService ===============
FROM golang:1.6

COPY ./src/login/ /login/

ENV GOPATH=/login/vendor/

WORKDIR /login/loginservice/

RUN mkdir /login/vendor && go get \
	&& CGO_ENABLED=0 go build -a -installsuffix cgo -v \
	-o loginservice.o LoginService.go

COPY loginservice.o /bin/loginservice

CMD ["bin/loginservice"]

