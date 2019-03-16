FROM golang:alpine

RUN apk update && apk upgrade && \
  apk add --no-cache git openssh curl vim bash gcc g++

RUN mkdir -p /go/src/app

ADD . /go/src/app

ENV GOPATH=/go

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/app

RUN dep ensure