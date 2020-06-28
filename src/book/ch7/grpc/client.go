package grpc

import (
	"book/ch7/pb"
	"context"
	"fmt"
	grpc1 "google.golang.org/grpc"
)

func main() {
	serviceAdd := "127.0.0.1:1234"
	conn,err   := grpc1.Dial(serviceAdd,grpc1.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	stringClient := pb.NewStringServiceClient(conn)
	stringReq    := &pb.StringRequest{A:"A",B:"B"}
	reply, _     := stringClient.Concat(context.Background(),stringReq)
	fmt.Println(reply.Ret)
}

