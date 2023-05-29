package main

import (
	"flag"

	"microservice-demo-go/pongservice/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "config file location")

// main
func main() {
	server.Run(*cfg)
}
