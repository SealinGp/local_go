package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

/*
通道,超时,计时器
1 s  = 1000 ms(毫秒)
1 ms = 1000 us(微妙)
1 us = 1000 ns(纳秒)
像圆周率π一样,e为数学中的常数, e = 2.71828182856904523536...
*/
//func init() {
//	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
//}
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
		"tcp1"  : tcp1,
		"tcp_client"  : tcp_client,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}
var clients = map[string]int{}
func tcp1()  {
	fmt.Println("starting the server ...")
	listener, err := net.Listen("tcp","127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error listening",err.Error())
		return
	}

	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("error reading",err.Error())
		}
	    go tcp_server(conn)
	}
}
func tcp_server(conn net.Conn)  {
	currentClient := ""
	for {
		buf := make([]byte,512)
		len1, err := conn.Read(buf)
		if err != nil {
			msg := "Error reading" + err.Error()
			if err == io.EOF {
				clients[currentClient] = 0
			}
			fmt.Println(msg)
			return
		}
		data := string(buf[:len1])

		//结束服务端
		if strings.Contains(data,"end server") {
			fmt.Println("end server")
			os.Exit(1)
		}

		//服务端名
		ix := strings.Index(data,"says")
		currentClient = data[0:ix-1]
		clients[currentClient] = 1

		//服务端列表记录
		if strings.Contains(data,"client list") {
			fmt.Println("this is the client list: 1:active 0:inactive")
			for name,ifActive := range clients {
				fmt.Println("User",name,"is",ifActive)
			}
		}
		fmt.Println("Received data",string(buf[:len1]))
	}
}
func tcp_client()  {
	conn,err := net.Dial("tcp","127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dialing",err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("what's your name ?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName,"\r\n")
	//给服务器发消息直到程序退出
	for {
		fmt.Println("send msg(type q to quit):")
		input,_ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input,"\r\n")
		if strings.ToUpper(trimmedInput) == "Q"{
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says " + trimmedInput))
	}
}
