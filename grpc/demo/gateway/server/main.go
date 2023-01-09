package main

import (
	"context"
	pb "demo002/proto/echo"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "the port to serve on")
var restful = flag.Int("restful", 8080, "the port to restful serve on")

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("client send message:%s\n", in.Message)
	return &pb.EchoResponse{Message: in.Message}, nil
}

func (s *server) ServerStreamingEcho(req *pb.EchoRequest, stream pb.Echo_ServerStreamingEchoServer) error {
	return nil
}
func (s *server) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	return nil
}
func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	return nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	log.Println("serving grpc on 0.0.0.0" + fmt.Sprintf(":%d", *port))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve:%v", err)
		}
	}()

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial server:%v", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterEchoHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *restful),
		Handler: gwmux,
	}
	log.Println("Serving gRPC-Gateway on http://0.0.0.0" + fmt.Sprintf(":%d", *restful))
	log.Fatalln(gwServer.ListenAndServe())
}
