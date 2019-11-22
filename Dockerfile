FROM golang:latest

WORKDIR /opt

RUN go get github.com/gorilla/websocket