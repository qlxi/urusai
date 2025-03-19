FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o urusai

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/urusai /app/
COPY config.json /app/

ENTRYPOINT ["/app/urusai"]
