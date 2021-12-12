FROM golang:1.14-alpine3.12 as builder

ENV BASE_APP_DIR /go/src/projekt
WORKDIR ${BASE_APP_DIR}

COPY . .

RUN go build -v -o projekt .

RUN mkdir /app && mv ./projekt /app/projekt

FROM alpine

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 1332
CMD ["/app/projekt"]