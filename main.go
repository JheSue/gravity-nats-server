package main

import (
	"log"
	"net"
	"net/url"
	"os"

	"github.com/nats-io/gnatsd/server"
)

func main() {
	//get container ip
	hostname, _ := os.Hostname()
	ipAddr, _ := net.ResolveIPAddr("ip", hostname)

	// generate routers object
	routes := []*url.URL{}
	routes = append(routes, &url.URL{
		Scheme: "nats",
		Host:   "nats-server-cluster-mgmt:6222",
	})

	// generate server options
	opts := server.Options{
		Host: "0.0.0.0",
		Port: 4222,
		Cluster: server.ClusterOpts{
			Host:      "0.0.0.0",
			Port:      6222,
			Advertise: ipAddr.String(),
		},
		Routes: routes,
	}

	// New server
	ser := server.New(&opts)
	ser.ConfigureLogger()

	// Run server
	err := server.Run(ser)
	if err != nil {
		log.Println(err)
	}

}
