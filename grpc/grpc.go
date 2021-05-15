package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-04-grpc.html
// 4.4.2 grpc 入门

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-05-grpc-hack.html
// 4.5.1 grpc 进阶
// 1.go get -u google.golang.org/grpc && go get -u github.com/golang/protobuf/protoc-gen-go
// 2.write hello.proto
// 3.protoc --go_out=plugins=grpc:. hello.proto
// 4.go run grpc.go grpc6 (grpc server)
// 5.go run grpc.go grpc7 (grpc client)
type HelloServiceImpl struct {
	*UnimplementedHelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello: " + args.GetValue()}
	return reply, nil
}

//grpc server
func Grpc6() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, &HelloServiceImpl{})

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server in :1234")
	grpcServer.Serve(lis)
}

//grpc client
func Grpc7() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-06-grpc-ext.html
// 4.6 grpc 和 protobuf 拓展
// 4.6.1 protobuf 验证器
// 4.6.2 grpc -> rest api
