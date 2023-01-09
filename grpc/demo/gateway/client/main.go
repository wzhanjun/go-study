package main

import (
	"context"
	"flag"
	"log"

	pb "demo002/proto/echo"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc dial failed:", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	r, err := client.UnaryEcho(context.Background(), &pb.EchoRequest{Message: "hello"})
	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}
	log.Printf("echo serve: %s", r.GetMessage())
}
