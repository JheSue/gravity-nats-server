module nats-server

go 1.20

require (
	github.com/nats-io/nats-server/v2 v2.10.11
	github.com/spf13/pflag v1.0.5
)

require (
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.5.5 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	go.uber.org/automaxprocs v1.5.3 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/time v0.5.0 // indirect
)

//replace github.com/nats-io/nats-server/v2 => ./test/nats-server
