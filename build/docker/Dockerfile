# build stage
FROM golang:1.15.12-alpine3.12 AS builder

WORKDIR /

COPY . .

RUN apk add --update build-base && apk upgrade --available
RUN go build -o /nats-server


# Final stage
FROM alpine:3.12.7

RUN apk update && apk upgrade --available

COPY --from=builder /nats-server /nats-server

EXPOSE 4222 6222

USER 1001

CMD ["/nats-server"]
