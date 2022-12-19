module nats-server

go 1.17

require (
	github.com/nats-io/nats-server/v2 v2.9.3
	github.com/spf13/pflag v1.0.5
)

require (
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.3.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	go.uber.org/automaxprocs v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20221010152910-d6f0a8c073c2 // indirect
	golang.org/x/sys v0.0.0-20221010170243-090e33056c14 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
)

//replace github.com/nats-io/nats-server/v2 => ./test/nats-server
