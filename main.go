package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	flag "github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

var (
	//defaultRoutes      = []string{"gravity-nats-0.gravity-nats-mgmt.default.svc.cluster.local"}
	defaultRoutes      = []string{}
	defaultClusterName = "nats-cluster"
	defaultHostname, _ = os.Hostname()
	defaultStoreDir    = "./datastore"
	defaultConfigFile  = ""
	routerMGMT         = flag.StringSlice("routes", defaultRoutes, "Set mgmt service name and port")
	clusterName        = flag.String("clusterName", defaultClusterName, "Set cluster name")
	serverName         = flag.String("serverName", defaultHostname, "Set server name")
	storeDir           = flag.String("storeDir", defaultStoreDir, "Set Store Dir")
	enableJetStream    = flag.Bool("enableJetStream", true, "enable JetStream")
	jetStreamMaxMemory = flag.Int64("jetStreamMaxMemory", 4000, "Set JetStream Max Memory, unit: MB")
	jetStreamMaxStore  = flag.Int64("jetStreamMaxStore", 0, "Set JetStream Max Store, unit: MB")
	enableMQTTBroker   = flag.Bool("enableMQTTBroker", false, "enable MQTT Broker")
	configFile         = flag.String("configFile", defaultConfigFile, "configuration file")
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
			Host:   r,
		})
	}

	sdir := fmt.Sprintf("%s/%s", *storeDir, *serverName)

	/*
		accounts := []*server.Account{}
		account := server.Account{}
		account.Name = "admin"
		account.Nkey = "admin"
		accounts = append(accounts, &account)

		users := []*server.User{}
		users = append(users, &server.User{
			Username: "admin",
			Password: "admin",
			Account:  &account,
			Permissions: &server.Permissions{
				Publish: &server.SubjectPermission{
					Allow: []string{">"},
				},
				Subscribe: &server.SubjectPermission{
					Allow: []string{">"},
				},
			},
		})
	*/
	opts := server.Options{
		Host:               "0.0.0.0",
		Port:               4222,
		HTTPHost:           "0.0.0.0",
		HTTPPort:           8222,
		MaxPayload:         1024 * 1024 * 64,
		WriteDeadline:      10 * time.Second,
		JetStream:          *enableJetStream,
		JetStreamMaxMemory: *jetStreamMaxMemory * 1024 * 1024,
		ServerName:         *serverName,
		StoreDir:           sdir,
		ConfigFile:         *configFile,
		//Users:         users,
		//Accounts:      accounts,
		//SystemAccount: "admin",
		//PidFile:  "/var/run/nats/nats.pid",
		//Debug:  true,
		//MaxConn:      1024 * 4,
		//SyncAlways:             true,
		SyncInterval:           30 * time.Second,
		NoSublistCache:         true,
		DisableJetStreamBanner: true,
		Logtime:                true,
	}
	if *jetStreamMaxStore != 0 {
		opts.JetStreamMaxStore = *jetStreamMaxStore * 1024 * 1024
	}

	if *configFile != "" {
		err := opts.ProcessConfigFile(*configFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Processing Cluster
	if len(*routerMGMT) > 0 {
		opts.Routes = routes
		opts.Cluster = server.ClusterOpts{
			Name:           *clusterName,
			Host:           "0.0.0.0",
			Port:           6222,
			Advertise:      ipAddr.String(),
			ConnectRetries: 600,
			PoolSize:       3,
			Compression: server.CompressionOpts{
				Mode: "s2_fast",
			},
		}
	}

	// Processing MQTT options
	if *enableMQTTBroker {
		mqttOpts := server.MQTTOpts{
			Host: "0.0.0.0",
			Port: 1883,
		}

		opts.MQTT = mqttOpts
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
		log.Fatal(err)
	}

	ser.WaitForShutdown()
}
