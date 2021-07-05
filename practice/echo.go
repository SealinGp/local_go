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
	"sync"
	"time"
)

//practice from https://books.studygolang.com/gopl-zh/ch8/ch8-04.html
func main() {
	var name string
	flag.StringVar(&name, "name", "", "...")
	flag.Parse()
	if name == "netcat2" {
		netcat2()
	}
	if name == "reverb1" {
		reverb1()
	}
	if name == "reverb2" {
		reverb2()
	}
	if name == "netcat3" {
		netcat3()
	}
}

func netcat3() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn := conn.(*net.TCPConn)

	done := make(chan struct{})

	//读server内容 -> output
	go func() {
		mustCopy1(os.Stdout, conn)
		log.Printf("done")
		done <- struct{}{}
	}()

	//写input内容 -> server
	mustCopy1(conn, os.Stdin)
	/*
		关闭网络连接中写方向的连接将导致server程序收到一个文件（end-of-file）结束的信号
		关闭网络连接中读方向的连接将导致后台goroutine的io.Copy函数调用返回一个“read from closed connection”（“从关闭的连接读”）类似的错误
		这里只关闭写,则代码块 [写input内容 -> server] 将会结束,而 [读server内容 -> output]将会继续,如果调用.Close(),关闭读和写,则都会退出
	*/
	tcpConn.CloseWrite()
	<-done
}

func netcat2() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//读取server内容 -> output
	go mustCopy1(os.Stdout, conn)

	//读input内容 -> server
	mustCopy1(conn, os.Stdin)
}

func mustCopy1(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}

func reverb1() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("[E] accept err %s", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), time.Second)
	}
}

func reverb2() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("[E] accept err %s", err)
			continue
		}
		go handleConn2(conn)
	}
}

func handleConn2(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo2(c, input.Text(), time.Second, wg)
	}
	wg.Wait()
	tcpC := c.(*net.TCPConn)
	tcpC.CloseWrite()
}

func echo2(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
