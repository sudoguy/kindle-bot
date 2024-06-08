FROM golang:latest as builder

COPY . /multistage

WORKDIR /multistage

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./app

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /multistage/bot .
