module nats-server

go 1.23.1

require (
	github.com/nats-io/nats-server/v2 v2.10.21
	github.com/spf13/pflag v1.0.5
	go.uber.org/automaxprocs v1.6.0
)

require (
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/nats-io/jwt/v2 v2.7.2 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/time v0.7.0 // indirect
)

//replace github.com/nats-io/nats-server/v2 => ./test/nats-server
