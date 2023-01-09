package main

import (
	"context"
	pb "demo002/helloworld/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("received:%v \n", in.GetName())
	return &pb.HelloReply{Message: "hello " + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})

	log.Println("serving grpc on 0.0.0.0" + port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

// ghz -c 2 -n 10  --insecure --proto ./proto/hello_world.proto --call helloworld.Greeter.SayHello -d {\"name\":\"Joe\"} 0.0.0.0:50051
// ghz -c 10 -n 1000 --insecure --proto ./proto/hello_world.proto --call helloworld.Greeter.SayHello -d {\"name\":\"Joe\"} --load-schedule=step --load-start=50 --load-step=10 --load-step-duration=5s -o report.html -O html 0.0.0.0:50051
