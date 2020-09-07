package main

import (
	"encoding/json"
	"github.com/nats-io/gnatsd/server"
	"log"
	"net/url"
	"os"
)

func main() {
	//get hostname
	hostname, _ := os.Hostname()

	// generate routers object
	routes := []*url.URL{}
	routes = append(routes, &url.URL{
		Scheme: "nats",
		Host:   "0.0.0.0:6222",
	})

	// generate server options
	opts := server.Options{
		Host: "0.0.0.0",
		Port: 4222,
		Cluster: server.ClusterOpts{
			ConnectRetries: -1,
			Host:           "0.0.0.0",
			Port:           6222,
			Advertise:      hostname,
		},
		Routes: routes,
	}
	optData, _ := json.Marshal(opts.Clone())

	log.Println(string(optData))

	// New server
	ser := server.New(&opts)

	log.Println(ser.NumRoutes())

	// Run server
	err := server.Run(ser)
	if err != nil {
		log.Println(err)
	}
}
