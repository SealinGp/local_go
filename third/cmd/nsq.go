package main

import "third/nsq"

func main() {
	nsq.Consume("test","channel")
}