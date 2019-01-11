package main

import (
	"grpc-test/server"
	"log"
	"os"
)

func main() {
	function := os.Getenv("FUNCTION")

	if "server" == function {
		server.StartServer()
	} else {
		log.Println("hello world")
	}
}
