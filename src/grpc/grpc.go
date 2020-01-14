package main

import (
	"log"
	"os"
)
//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch4-rpc/ch4-01-rpc-intro.md

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required")
	}
	f := map[string]func(){
		"grpc1":grpc1,
	}
	f[os.Args[1]]()
}

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch4-rpc/ch4-02-pb-intro.md
type HelloService struct {}
func (h *HelloService)Hello(request *String,reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
func grpc1()  {

}