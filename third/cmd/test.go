package main

import (
	"flag"
	"third/test"
)

var server int

func main() {
	flag.IntVar(&server,"server",0,"https server(1) or client(0)")
	flag.Parse()

	if server == 1 {
		test.HttpsServer("../test/client.crt","../test/server.crt","../test/server.key")
		return
	}
	test.HttpsClient("../test/server.crt","../test/client.crt","../test/client.key")
}