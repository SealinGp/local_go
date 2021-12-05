package main

import (
	"github.com/SealinGp/local_go/grpc"
)

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch4-rpc/ch4-01-rpc-intro.md

var grpcFuncs = map[string]func(){
	"grpc6": grpc.Grpc6,
	"grpc7": grpc.Grpc7,
}
