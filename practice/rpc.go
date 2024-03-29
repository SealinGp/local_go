package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.9.md
适用于两服务端均为 基于go 的 RPCs服务 之间的数据传输
基于 gob.go里面的 gob 包
用于在两台不同的服务器之间go服务的调用

注:有问题,客户端无法调用,已提交github issue,待更新
*/

var rpcFuns = map[string]func(){
	"rpc1": rpc1,
	"rpc2": rpc2,
	"rpc3": rpc3,
}

//server
func rpc1() {
	//rpc register
	cal := new(Args)
	_ = rpc.Register(cal)
	rpc.HandleHTTP()

	//start server
	e := http.ListenAndServe(":1234", rpc.DefaultServer)
	if e != nil {
		log.Println(e)
	}
}

type Args struct {
	N, M int
}

func (*Args) Multiply(a Args, reply *int) error {
	*reply = a.M * a.N
	return nil
}

//client
func rpc2() {
	//connect
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("dial rpc server(localhost:1234) failed:" + err.Error())
		return
	}
	fmt.Println("connect success! start to call")

	//call
	defer client.Close()
	arg := Args{2, 3}
	reply := 0
	err = client.Call("Args.Multiply", arg, &reply)
	if err != nil {
		fmt.Println("cal error:" + err.Error())
		return
	}
	fmt.Println("call success,ready to print result")

	fmt.Println(reply)
}

func rpc3() {
	a := &Args{2, 3}
	r := 0
	a.Multiply(*a, &r)
	fmt.Println(r)
}
