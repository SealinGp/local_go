package grpcTest

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"math/rand"
	"net"
	"runtime"
	"strings"
	"third/grpcTest/hello"
	"time"
)

var (
	grpcAddr = "localhost:8443"
)

type server struct {
	hello.UnimplementedHelloServiceServer
}

func (s *server)Hello(ctx context.Context,helloStringReq *hello.HelloString) (*hello.HelloString,error) {
	//sleep 1 year
	//time.Sleep(time.Hour*12*30*12)

	helloStringResp := &hello.HelloString{
		Value: "received:" + helloStringReq.Value,
	}
	return helloStringResp,nil
}

func GrpcServer()  {
	//TLS证书配置
	cert,err := tls.LoadX509KeyPair("../grpcTest/public.crt","../grpcTest/private.key")
	if err != nil {
		log.Fatal("credentials err:",err)
	}
	transportCre := credentials.NewServerTLSFromCert(&cert)
	grpcOptions := []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Second*5),
		grpc.Creds(transportCre),
	}

	//开始监听
	listener,err := net.Listen("tcp",grpcAddr)
	if err != nil {
		log.Fatal("failed listen:",err)
	}
	log.Println("start listen",grpcAddr,"...")

	//注册服务,并处理连接
	grpcServer := grpc.NewServer(grpcOptions...)
	hello.RegisterHelloServiceServer(grpcServer,&server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GrpcClient()  {
	//开始连接
	transportCred,err := credentials.NewClientTLSFromFile("../grpcTest/public.crt","localhost")
	if err != nil {
		log.Fatal("credentials err:",err)
	}
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCred),
	}
	conn,err := grpc.Dial(grpcAddr,dialOptions...)
	if err != nil {
		log.Fatal("dial err:",err)
	}
	defer conn.Close()


	taskCh := make(chan func(),1000)
	for i := 0; i < 500; i++ {
		go func() {
			for task := range taskCh {
				task()
			}
		}()
	}
	go func() {
		for range time.Tick(time.Second) {
			log.Println("goroutine num:",runtime.NumGoroutine())
		}
	}()

	//开始请求
	rand.Seed(time.Now().UnixNano())
	input := 1
	for range time.Tick(6 * time.Second) {
		//fmt.Scanln(&input)
		if input > 1 {
			input = 2000
		}

		reqParamVal := strings.Repeat(string('z' - rand.Intn(25)),1024)
		for i := 0; i < input; i++ {
			taskCh <- func() {
				helloClient := hello.NewHelloServiceClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				_,err := helloClient.Hello(ctx,&hello.HelloString{Value:reqParamVal})
				if err != nil {
					log.Println("invoke err",err)
					return
				}
				//log.Println("received",resp.Value)
			}
		}
		input++
	}
}