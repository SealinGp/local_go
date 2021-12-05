package main

import (
	"log"

	"github.com/SealinGp/local_go/third/quic"
)

func QuicFunc() {
	go func() {
		log.Fatal(quic.EchoServer())
	}()

	err := quic.ClientMain()
	if err != nil {
		panic(err)
	}
}
