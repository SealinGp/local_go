package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	
}

func clientWrite()  {
	influxdb2.NewClient("http://localhost:8086","token")
}