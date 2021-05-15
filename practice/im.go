package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

//https://books.studygolang.com/gopl-zh/ch8/ch8-10.html

func main() {
	var name string
	flag.StringVar(&name, "name", "", "")
	flag.Parse()
	if name == "server" {
		IMServer()
	}
	if name == "client" {
		IMClient()
	}
}

func IMServer() {
	lis, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Printf("[E] listen err:%s", err)
		return
	}
	defer lis.Close()

	go broadcaster()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("[E] accept err:%s", err)
			continue
		}
		go IMServerHandleConn(conn)
	}
}

func IMServerHandleConn(c net.Conn) {
	ch := make(chan string)
	defer c.Close()
	go clientWriter(c, ch)

	who := c.RemoteAddr().String()

	//欢迎语发送到该连接的客户端中
	ch <- fmt.Sprintf("You are %s", who)

	//该用户进入房间
	entering <- Client{
		cli:  ch,
		addr: who,
	}
	//进入房间提示语
	messages <- fmt.Sprintf("%s has come to the house", who)

	//从该连接中读取数据
	input := bufio.NewScanner(c)
	for input.Scan() {
		text := input.Text()

		//离开房间的消息
		if strings.Contains(text, "exit") {
			break
		}

		//将该连接中读取到的数据发送给每个客户端
		messages <- who + ": " + text
	}

	//离开房间?
	leaving <- ch
	messages <- who + " has left"
}

//把从客户端接收到的消息写入该tcp连接中
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func IMClient() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Printf("[E] dial err:%s \n", err)
		return
	}
	defer conn.Close()

	//从连接读取数据
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println(err)
		}
	}()

	//把数据写入连接
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()

		fmt.Fprintln(conn, text)
		if err != nil {
			log.Printf("[E] conn write err:%s \n", err)
		}
		if strings.Contains(text, "exit") {
			break
		}
	}
}

type client chan<- string
type Client struct {
	cli  client
	addr string
}

var (
	entering = make(chan Client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]string)
	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				return
			}
			//往每个客户端写入该用户发送的消息
			for cli, addr := range clients {
				//不给自己发消息
				if strings.Contains(msg, addr) {
					continue
				}
				cli <- msg
			}
		case cli := <-entering:
			clients[cli.cli] = cli.addr
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
