#=============== LoginService Build Image ===============
FROM google/golang:1.6

ADD . /src
WORKDIR 

ENV GOPATH=

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -o loginservice.o src/login/loginservice/LoginService.go
EXPORT loginservice.o

#=============== LoginService ===============
FROM busybox
IMPORT loginservice.o /bin/loginservice

CMD ["bin/loginservice"]

TAG loginservice:latest
