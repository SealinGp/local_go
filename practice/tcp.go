package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"syscall"
)

/*
ref :https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.1.md
通道,超时,计时器
1 s  = 1000 ms(毫秒)
1 ms = 1000 us(微妙)
1 us = 1000 ns(纳秒)
像圆周率π一样,e为数学中的常数, e = 2.71828182856904523536...
*/
//func init() {
//	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
//}
var tcpFuncs = map[string]func(){
	"tcp1":       tcp1,
	"tcp_client": tcp_client,
	"tcp2":       tcp2,
}
var clients = map[string]int{}

//类似于socket编程,客户端,服务端消息传送(=聊天系统后台原理)
func tcp1() {
	fmt.Println("starting the server ...")
	listener, err := net.Listen("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error reading", err.Error())
		}
		go tcp_server(conn)
	}
}
func tcp_server(conn net.Conn) {
	currentClient := ""
	for {
		buf := make([]byte, 512)
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
		if strings.Contains(data, "end server") {
			fmt.Println("end server")
			os.Exit(1)
		}

		//服务端名
		ix := strings.Index(data, "says")
		currentClient = data[0 : ix-1]
		clients[currentClient] = 1

		//服务端列表记录
		if strings.Contains(data, "client list") {
			fmt.Println("this is the client list: 1:active 0:inactive")
			for name, ifActive := range clients {
				fmt.Println("User", name, "is", ifActive)
			}
		}
		fmt.Println("Received data", string(buf[:len1]))
	}
}
func tcp_client() {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("what's your name ?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	//给服务器发消息直到程序退出
	for {
		fmt.Println("send msg(type q to quit):")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if strings.ToUpper(trimmedInput) == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says " + trimmedInput))
	}
}

//改进版本的socket
const readBuffer = 25 //读取数据缓冲大小
func tcp2() {
	flag.Parse()
	if flag.NArg() != 3 {
		panic("command: go build tcp.go && tcp tcp2 host port")
	}

	hostPort := net.JoinHostPort(flag.Arg(1), flag.Arg(2))
	listener := initServer(hostPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept failed")
		connHandler(conn)
	}
}

//服务端初始化(创建,监听,返回)
func initServer(hostPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostPort)
	checkError(err, "resolve failed,address:"+hostPort)
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "listen failed,address:"+hostPort)
	fmt.Println("listening to ", hostPort)
	return listener
}
func checkError(err error, info string) {
	if err != nil {
		panic("ERROR:" + info + " " + err.Error())
	}
}
func connHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	fmt.Println("connect from", connFrom)
	for {
		var ibuf []byte = make([]byte, readBuffer+1)
		length, err := conn.Read(ibuf[0:readBuffer])
		ibuf[readBuffer] = 0 //防止溢出?
		switch err {
		case nil:
			printMsg(length, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	fmt.Println("closed connection", connFrom)
	checkError(err, "Close error")
}
func printMsg(length int, err error, ibuf []byte) {
	if length > 0 {
		fmt.Println("<", length, ":")
		fmt.Println(string(ibuf))
		fmt.Println(">")
	}
}
