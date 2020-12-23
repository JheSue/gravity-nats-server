# build stage
FROM golang:1.15.3-alpine3.12 AS builder

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 go build -installsuffix cgo -o /nats-server


# Final stage
FROM alpine:3.12

RUN addgroup -S gravity && adduser -S gravity -G gravity
COPY --from=builder /nats-server /nats-server

EXPOSE 4222 6222

USER gravity
CMD ["/nats-server"]
