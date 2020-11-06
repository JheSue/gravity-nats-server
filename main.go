package main

import (
	"log"
	"net"
	"net/url"
	"os"

	"flag"
	"github.com/nats-io/nats-server/v2/server"
)

var routerMGMT = flag.String("routes", "nats-server-cluster-mgmt:6222", "Input mgmt service name and port")

func main() {
	flag.Parse()
	if *routerMGMT == "" {
		log.Fatal("--routes are required.")
	}

	//get container ip
	hostname, _ := os.Hostname()
	ipAddr, _ := net.ResolveIPAddr("ip", hostname)

	// generate routers object
	routes := []*url.URL{}
	routes = append(routes, &url.URL{
		Scheme: "nats-route",
		//Host:   "nats-server-cluster-mgmt:6222",
		Host: *routerMGMT,
	})

	// generate server options
	opts := server.Options{
		Host: "0.0.0.0",
		Port: 4222,
		Cluster: server.ClusterOpts{
			Host:           "0.0.0.0",
			Port:           6222,
			Advertise:      ipAddr.String(),
			ConnectRetries: 10,
		},
		Routes: routes,
		//Debug:  true,
		HTTPHost: "0.0.0.0",
		HTTPPort: 8222,
		//PidFile:  "/var/run/nats/nats.pid",
	}

	// New server
	ser, err := server.NewServer(&opts)
	if err != nil {
		log.Fatal(err)
	}
	ser.ConfigureLogger()

	// Run server
	err = server.Run(ser)
	if err != nil {
		log.Println(err)
	}

}
