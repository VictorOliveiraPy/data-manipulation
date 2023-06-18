FROM golang:1.20-alpine

ADD . /app

WORKDIR /app

RUN go get -d -v ./...
RUN go install -v ./...