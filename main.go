package main

import (
	"log"
	"net"
	"net/url"
	"os"

	"github.com/nats-io/nats-server/v2/server"
	flag "github.com/spf13/pflag"
)

var (
	defaultRoutes = []string{"gravity-nats-0.gravity-nats-mgmt.default.svc.cluster.local"}
	routerMGMT    = flag.StringSlice("routes", defaultRoutes, "Input mgmt service name and port")
)

func main() {
	flag.Parse()

	//get container ip
	hostname, _ := os.Hostname()
	ipAddr, _ := net.ResolveIPAddr("ip", hostname)

	// generate routes object
	routes := []*url.URL{}
	for _, r := range *routerMGMT {
		routes = append(routes, &url.URL{
			Scheme: "nats-route",
			//Host:   "nats-server-cluster-mgmt:6222",
			Host: r,
		})
	}

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
		//Debug:  true,
		Routes:     routes,
		HTTPHost:   "0.0.0.0",
		HTTPPort:   8222,
		MaxPayload: 1024 * 1024 * 10,
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
