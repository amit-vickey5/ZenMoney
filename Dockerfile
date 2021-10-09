FROM golang:1.17.2-alpine3.13
RUN mkdir /build
ADD go.mod go.sum server.go /build/
WORKDIR /build
RUN go build