FROM golang:1.16.2-alpine as builder

RUN apk update && apk add --no-cache make gcc libc-dev protobuf-dev
RUN go get github.com/golang/protobuf/protoc-gen-go

ENV GO11MODULE=on
ENV MODE=Production
ENV WEBSITE_ORIGIN=http://localhost:3000
ENV WEBSITE_DOMAIN=localhost
ENV PORT=8080
ENV ACCESS_TOKEN_LIFETIME=8h
ENV DATA_MANAGEMENT_CONNECTION=5g-v2x-data-management-service:8082
ENV USER_CONNECTION=5g-v2x-user-service:8083


WORKDIR /app
ADD Makefile /app/Makefile
ADD third_party /app/third_party
ADD cmd /app/cmd
ADD pkg /app/pkg
ADD go.mod /app/go.mod
ADD go.sum /app/go.sum
ADD api /app/api
ADD internal /app/internal

RUN go mod tidy
RUN go mod vendor

RUN make proto
RUN make bin

CMD [ "/app/bin/server-linux-amd64" ]

EXPOSE 8080