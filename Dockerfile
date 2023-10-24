FROM golang:1.21.1-alpine3.18

RUN mkdir /go-postgres

COPY . /go-postgres

WORKDIR /go-postgres

RUN go build -o goapp .

EXPOSE 8080

CMD ["/go-postgres/goapp"]