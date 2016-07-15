#=============== LoginService ===============
FROM golang:1.6

COPY ./src/ /src/

ENV GOPATH=/src/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -o loginservice.o /src/login/loginservice/LoginService.go

COPY loginservice.o /bin/loginservice

CMD ["bin/loginservice"]

