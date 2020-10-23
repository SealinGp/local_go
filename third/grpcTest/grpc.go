package grpcTest

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"third/grpcTest/hello"
	"time"
)

var (
	grpcAddr = "localhost:8443"
	connectionTimeout = time.Second*5
)

type server struct {
	hello.UnimplementedHelloServiceServer
}

func (s *server)Hello(ctx context.Context,helloStringReq *hello.HelloString) (*hello.HelloString,error) {
	helloStringResp := &hello.HelloString{
		Value: "received:" + helloStringReq.Value,
	}
	return helloStringResp,nil
}

func GrpcServer()  {
	//开始监听
	listener,err := net.Listen("tcp",grpcAddr)
	if err != nil {
		log.Fatal("failed listen:",err)
	}
	log.Println("start listen",grpcAddr,"...")

	//注册服务,并处理连接

	grpcOptions := []grpc.ServerOption{
		grpc.ConnectionTimeout(connectionTimeout),
	}
	grpcServer := grpc.NewServer(grpcOptions...)
	hello.RegisterHelloServiceServer(grpcServer,&server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GrpcClient()  {
	//开始连接
	conn,err := grpc.Dial(grpcAddr,grpc.WithInsecure())
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

	//开始请求
	input := 1
	for {
		fmt.Scanln(&input)
		if input == -1 {
			return
		}

		for i := 0; i < input; i++ {
			taskCh <- func() {
				helloClient := hello.NewHelloServiceClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				resp,err := helloClient.Hello(ctx,&hello.HelloString{Value:"client say hello"})
				if err != nil {
					log.Println("invoke err",err)
					return
				}
				log.Println("received",resp.Value)
			}
		}
	}
}