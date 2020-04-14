FROM golang:stretch as builder

COPY . /multistage

WORKDIR /multistage

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /multistage .
CMD ["./bot"]