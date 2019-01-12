package main

import (
	"grpc-test/client"
	"grpc-test/server"
	"os"
)

func main() {
	function := os.Getenv("FUNCTION")

	if "server" == function {
		server.StartServer()
	} else {
		// client.GetFeature()
		client.StartClientServer()
	}
}
