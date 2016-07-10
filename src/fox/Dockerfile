#=============== FoxService Build Image ===============
FROM google/golang:1.6

ADD . /src
WORKDIR 

ENV GOPATH=

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -o foxservice.o src/fox/foxservice/FoxService.go
EXPORT foxservice.o

#=============== FoxService ===============
FROM busybox
IMPORT foxservice.o /bin/foxservice

CMD ["bin/foxservice"]

TAG foxservice:latest




