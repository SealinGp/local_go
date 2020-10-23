package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

/**
问题描述:
1.服务端tcp接收缓冲区/客户端tcp发送缓冲区满时,客户端会卡住,超时设置将不起作用
go version 1.13.5 客户端:go https client 服务端: go https(2) server
在使用使用https 2.x 的情况下,当服务端卡住(ctrl + z｜ dlv可以停止该服务)的情况下,
客户端在同一条tcp连接上不断请求直到 服务端的tcp接收缓冲区 和 客户端的发送缓冲区 被塞满的时候,
golang https client的请求超时设置将不起作用而导致客户端卡住
问题分析: http2 使用了多路复用,一条tcp连接多个请求复用,若在客户端配置 不强制使用http2请求即可暂时解决

2.服务端卡死,
2个服务端的情况下(ip不同,DNS解析域名得出来的ip会有两个服务),其中一个服务端f

server
./test -server=true

client with http2(default)
./test

client without http2
./test -http2=false
*/

var (
	addr = "localhost:8443"
	//服务端超时时间
	//var serverTimeout = time.Second*2
	//客户端超时时间
	clientTimeout = time.Second * 2
	//服务端处理时间时间
	handlerTime = time.Hour*24*30*12

	forceHttp2 bool
	Server     bool
)

func main() {
	flag.BoolVar(&Server, "server", false, "https server(true) or client(false)")
	flag.BoolVar(&forceHttp2, "http2", true, "client if force http2")
	flag.Parse()

	if Server {
		HttpsServer("client.crt", "server.crt", "server.key")
		return
	}
	HttpsClient("server.crt", "client.crt", "client.key")
}

type handler struct{}
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	time.Sleep(handlerTime)

	n, err := w.Write([]byte("req ok"))
	if err != nil {
		log.Println("write err:", err)
		return
	}
	log.Println("write ok!", n, "bytes")
	w.WriteHeader(http.StatusOK)

}

//https server
func HttpsServer(clientCrt, serverCrt, serverKey string) {
	//证书读取配置
	caCert, err := ioutil.ReadFile(clientCrt)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}



	//开始监听
	httpSvr := &http.Server{
		Addr:      addr,
		Handler:   &handler{},
		TLSConfig: cfg,
		WriteTimeout:time.Second*2,
		ReadTimeout:time.Second*2,
		IdleTimeout:time.Second*10,
	}
	log.Printf("listen in https://%s , pid=%d ... \n", addr, os.Getpid())
	log.Fatal(httpSvr.ListenAndServeTLS(serverCrt, serverKey))

}

//https client
func HttpsClient(serverCrt, clientCrt, clientKey string) {
	//客户端证书配置
	caCert, err := ioutil.ReadFile(serverCrt)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     forceHttp2,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,

			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{cert},
			},
		},
		Timeout: clientTimeout,
	}

	//准备请求
	rand.Seed(time.Now().UnixNano())
	taskCh := make(chan func(), 1000)
	for i := 0; i < 500; i++ {
		go func() {
			for task := range taskCh {
				task()
			}
		}()
	}
	input := 1
	for _ = range time.Tick(time.Second*2) {
		//获取键盘输入作为请求次数,默认为1
		//input := 1
		//fmt.Println("please input:")
		//fmt.Scanln(&input)
		if input > 1 {
			input = 100
		}
		log.Printf("req https://%s timeout=%s input=%d \n", addr, clientTimeout, input)

		//开始请求
		for i := 0; i < input; i++ {
			taskCh <- func() {
				//给body随机数防止tcp发送缓冲区压缩
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				data := make([]byte, 0, 1024)
				buf := bytes.NewBuffer(data)
				for j := 0; j < cap(data)/8; j++ {
					err := binary.Write(buf, binary.BigEndian, rand.Uint64())
					if err != nil {
						log.Println(err)
						return
					}
				}

				req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("https://%s", addr), buf)
				if err != nil {
					log.Println(err)
					return
				}

				resp, err := client.Do(req)
				if err != nil {
					log.Println(err)
					return
				}
				defer resp.Body.Close()

				//获取数据
				data, err = ioutil.ReadAll(resp.Body)
				log.Printf("received status:%s body:%s \n", resp.Status, string(data))
			}
		}
		input++
	}
}
