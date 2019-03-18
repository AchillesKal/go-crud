FROM golang:alpine

RUN apk update && apk upgrade && \
  apk add --no-cache git openssh curl vim bash gcc g++

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY . .

ENV GOPATH=/go

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/cache
WORKDIR /go/src/cache

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

ENTRYPOINT [ "/go/src/app/docker/entrypoint.sh" ]

WORKDIR /go/src/app

