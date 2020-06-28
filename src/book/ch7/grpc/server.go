package grpc

import (
	string_service "book/ch7/grpc/string-service"
	"book/ch7/pb"
	"flag"
	"log"
	"net"
	grpc1 "google.golang.org/grpc"
)

func main()  {
	flag.Parse()
	lis, err := net.Listen("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen:%v",err)
	}
	grpcServer := grpc1.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer,stringService)
	grpcServer.Serve(lis)
}