FROM golang:1.14-alpine

RUN apk add build-base vim git
RUN go get -v github.com/rubenv/sql-migrate/...
RUN go get github.com/pilu/fresh

ADD . /go/src/app

WORKDIR /go/src/app

RUN go mod download