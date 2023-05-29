package main

import (
	"flag"

	"microservice-demo-go/pingservice/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "config file location")

// main
func main() {
	server.Run(*cfg)
}
