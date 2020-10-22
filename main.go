package main

import (
	"log"
	"net"
	"net/url"
	"os"

	"flag"
	"github.com/nats-io/gnatsd/server"
)

var routerMGMT = flag.String("routers", "nats-server-cluster-mgmt:6222", "Input mgmt service name and port")

func main() {
	flag.Parse()
	if *routerMGMT == "" {
		log.Fatal("--routers are required.")
	}

	//get container ip
	hostname, _ := os.Hostname()
	ipAddr, _ := net.ResolveIPAddr("ip", hostname)

	// generate routers object
	routes := []*url.URL{}
	routes = append(routes, &url.URL{
		Scheme: "nats",
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
