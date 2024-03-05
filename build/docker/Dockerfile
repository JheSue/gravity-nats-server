# build stage
FROM golang:1.21.7-alpine3.19 AS builder

WORKDIR /

COPY . .

RUN apk add --update build-base && apk upgrade --available
RUN go build -o /nats-server


# Final stage
FROM alpine:3.19.1

WORKDIR /

RUN apk update && apk upgrade --available && apk add tzdata
RUN adduser -u 1001 1001 -D

COPY --from=builder /nats-server /nats-server
COPY --from=builder /build/docker/generate-key.sh /generate-key.sh
COPY --from=natsio/nats-box:latest /usr/local/bin/nats /usr/local/bin/nsc /usr/local/bin/

RUN mkdir -p  /datastore && \
	chown -R 1001:0 /datastore /generate-key.sh && \
	chmod -R g+rwX /datastore && \
	chmod 750 /generate-key.sh

USER 1001

ENV NKEYS_PATH /.nsc/nkeys
ENV XDG_DATA_HOME /.nsc
ENV XDG_CONFIG_HOME /.nsc/.config

EXPOSE 4222 6222 1883
ENV TZ="Asia/Taipei"

CMD ["/nats-server"]
