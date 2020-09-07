# build stage
FROM golang:1.13-alpine AS builder

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 go build -installsuffix cgo -o /nats-server


# Final stage
FROM alpine:3.10

COPY --from=builder /nats-server /nats-erver

EXPOSE 4222 6222

CMD ["/nats-server"]
