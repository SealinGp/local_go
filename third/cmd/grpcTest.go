package main

import (
	"flag"
	"third/grpcTest"
)

var (


	IsServer     bool
)

func main() {
	flag.BoolVar(&IsServer, "server", false, "https server(true) or client(false)")
	flag.Parse()

	if IsServer {
		grpcTest.GrpcServer()
		return
	}
	grpcTest.GrpcClient()
}