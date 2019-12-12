package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.9.md
适用于两服务端均为 基于go 的 RPCs服务 之间的数据传输
基于 gob.go里面的 gob 包
用于在两台不同的服务器之间go服务的调用

注:有问题,客户端无法调用,已提交github issue,待更新
*/
func init() {
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}
func execute(n string) {
	funs := map[string]func(){
		"rpc1" : rpc1,
		"rpc2" : rpc2,
		"rpc3" : rpc3,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//server
func rpc1()  {
	//rpc register
	cal := new(Args)
	_ = rpc.Register(cal)
	rpc.HandleHTTP()

	//start server
	lis,err := net.Listen("tcp","localhost:1234")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("rpc server listening in: localhost:1234")
	go func() {
		err := http.Serve(lis,nil)
		if err != nil {
			fmt.Println("serve err:",err.Error())
		}
	}()
	time.Sleep(time.Second*1)
	select {}
}
type Args struct {
	N,M int
}
func (*Args)Multiply(a Args,reply *int) error {
	*reply = a.M * a.N
	return nil
}

//client
func rpc2()  {
	//connect
	client,err := rpc.DialHTTP("tcp","localhost:1234")
	if err != nil {
		fmt.Println("dial rpc server(localhost:1234) failed:"+err.Error())
		return
	}
	fmt.Println("connect success! start to call")

	//call
	defer client.Close()
	arg     := Args{2,3}
	reply   := 0
	err     = client.Call("Args.Multiply",arg,&reply);
	if err != nil {
		fmt.Println("cal error:"+err.Error())
		return
	}
	fmt.Println("call success,ready to print result")

	fmt.Println(reply)
}

func rpc3()  {
	a := &Args{2,3}
	r := 0
	a.Multiply(*a,&r)
	fmt.Println(r)
}