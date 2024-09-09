# build stage
FROM golang:1.23.1-alpine3.20 AS builder

ENV GOBIN=/

ARG VERSION_NSC=2.8.7
ARG VERSION_NATS=0.1.5

WORKDIR /

COPY . .

RUN apk add --update build-base upx && apk upgrade --available
RUN go build -ldflags="-s -w" -o /nats-server

RUN go install -ldflags="-X main.version=${VERSION_NSC} -s -w" github.com/nats-io/nsc/v2@v${VERSION_NSC} &&\
	go install -ldflags="-s -w" github.com/nats-io/natscli/nats@v${VERSION_NATS}

RUN upx -6 /nats /nsc /nats-server

# Final stage
FROM alpine:3.20

WORKDIR /

RUN apk update && apk upgrade --available && apk add tzdata
RUN adduser -u 1001 1001 -D

COPY --from=builder /nats-server /nats-server
COPY --from=builder /build/docker/generate-key.sh /generate-key.sh
COPY --from=builder /nats /nats
COPY --from=builder /nsc /nsc
#COPY --from=natsio/nats-box:latest /usr/local/bin/nats /usr/local/bin/nsc /usr/local/bin/

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
