FROM golang:alpine

RUN apk add build-base vim git
RUN go get github.com/pilu/fresh

ADD . /go/src/app

WORKDIR /go/src/app

RUN go mod download