FROM golang:1.16-alpine AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0


COPY . /app
WORKDIR /app

RUN go build echoserver.go

FROM alpine:latest

WORKDIR /

COPY --from=build /app .

USER 1001
ENTRYPOINT ./echoserver
