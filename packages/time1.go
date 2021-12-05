package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

var (
	port string
	zone int
)

func time11() {
	flag.StringVar(&port, "port", "", "")
	flag.IntVar(&zone, "zone", 0, "")
	flag.Parse()
	if port == "" {
		log.Fatal("[E] port required")
		return
	}

	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}
func handleConn(c net.Conn) {

	defer c.Close()
	for {
		now := time.Now()
		if zone > 0 {
			now = now.Add(time.Hour * 24 * time.Duration(zone))
		}
		_, err := io.WriteString(c, now.Format(time.Stamp+"\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

type info struct {
	Time string
	Port string
}

func clockwall() {
	var (
		ports string
	)
	flag.StringVar(&ports, "ports", "", "")
	flag.Parse()
	if ports == "" {
		log.Fatal("[E] ports required")
	}
	portsSli := strings.Split(ports, ",")
	portsLen := len(portsSli)

	outputCh := make(chan info, portsLen)
	for _, port := range portsSli {
		go Dial(port, outputCh)
	}
	for {
		outputs := make([]info, 0, portsLen)
		for i := 0; i < portsLen; i++ {
			outputs = append(outputs, <-outputCh)
		}
		v, _ := json.Marshal(outputs)
		fmt.Println(string(v))
	}

}

func Dial(port string, outputCh chan info) {
	conn, err := net.Dial("tcp", net.JoinHostPort("localhost", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		output := make([]byte, 1024)
		n, err := conn.Read(output)
		if err != nil {
			log.Fatal(err)
		}
		outputCh <- info{
			Port: port,
			Time: string(output[:n]),
		}
	}
}
