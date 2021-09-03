package main

import (
	"context"
	"log"
	"net"

	pb "helloworld/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

func main() {

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建grpc 服务容器
	server := grpc.NewServer()
	pb.RegisterHelloworldServer(server, new(HelloworldService))

	log.Printf("server listening at %v", listen.Addr())

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type HelloworldService struct {
}

func (h *HelloworldService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}
