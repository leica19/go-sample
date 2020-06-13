FROM golang:alpine

RUN apk add build-base vim git

ADD . /go/src/app

WORKDIR /go/src/app

RUN go mod download