#=============== FoxService ===============
FROM golang:1.6

COPY ./src/ /src/

ENV GOPATH=/src/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -v -o foxservice.o /src/fox/foxservice/FoxService.go

COPY foxservice.o /bin/foxservice

CMD ["bin/foxservice"]

