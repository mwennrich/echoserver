FROM golang:1.17-alpine AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0


COPY . /app
WORKDIR /app

RUN go build echoserver.go
RUN strip echoserver

FROM alpine:latest

WORKDIR /

COPY --from=build /app .

USER 1001
ENTRYPOINT ./echoserver
