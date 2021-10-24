FROM golang:1.14-alpine3.12 as builder

ENV BASE_APP_DIR /go/src/github.com/projektystudenckie/finance_restapi
WORKDIR ${BASE_APP_DIR}

COPY . .

RUN go build -v -o finance_restapi .

FROM alpine

WORKDIR /app