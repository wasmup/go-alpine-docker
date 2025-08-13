FROM alpine:latest

WORKDIR /usr/local

COPY go ./go

ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH
