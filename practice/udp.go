package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	isServer = flag.Bool("s", true, "server ?")
	address  = flag.String("addr", "localhost:1234", "server or client address")
	sigCh    = make(chan os.Signal)
	stopCh   = make(chan bool)
)

func init() {
	flag.Parse()
}

//server: go run udp.go -s=true -addr="127.0.0.1:1234"
//client: go run udp.go -s=false -addr="127.0.0.1:1234"
func main() {
	log.Println(*isServer)

	if *isServer {
		go udpServer()
		notifySignal()
		return
	}
	go udpClient()
	notifySignal()
}

func udpServer() {
	serverAddr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp", serverAddr)

	defer conn.Close()
	for {
		select {
		case <-stopCh:
			break
		default:
		}

		//read
		input := make([]byte, 1500)
		n, remoteAddr, err := conn.ReadFromUDP(input)
		if err != nil {
			log.Printf("read err:%s", err)
			return
		}
		log.Printf("received raddr:%s msg:%s", remoteAddr, input[:n])
	}
}

func udpClient() {
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	//keyboard scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-stopCh:
			break
		default:
		}

		//write
		raddr, _ := net.ResolveUDPAddr("udp", *address)
		_, err := conn.WriteToUDP(scanner.Bytes(), raddr)
		if err != nil {
			return
		}
	}
}

func notifySignal() {
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-sigCh:
		close(stopCh)
		log.Printf("received sig %s \n", sig)
		return
	}
}
