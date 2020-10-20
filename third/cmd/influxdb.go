package main

import (
	"context"
	"log"
	"os"
	"third/influxdb"
)

func main() {
	log.Println()
	client := influxdb.NewInfluxDbClient("http://localhost:8086","",log.New(os.Stderr,"[influxdb]",log.LstdFlags),context.Background())
	defer client.Close()

	//client.Write(true)
	client.Query()
}
