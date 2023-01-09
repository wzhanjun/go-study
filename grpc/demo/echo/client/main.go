package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "demo002/echo/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	// unary(client)
	// serverStream(client)
	// clientStream(client)
	bidirectionalSteram(client)
}

func unary(client pb.EchoClient) {
	resp, err := client.UnaryEcho(context.Background(), &pb.EchoRequest{Message: "unary request"})
	if err != nil {
		log.Printf("send error: %v \n", err)
	}
	fmt.Printf("received :%v \n", resp.GetMessage())
}

func serverStream(client pb.EchoClient) {
	stream, err := client.ServerStreamingEcho(context.Background(), &pb.EchoRequest{Message: "hello world"})
	if err != nil {
		log.Fatalf("could not echo :%+v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("receiver error:%v", err)
			continue
		}
		log.Printf("received data:%v", resp.GetMessage())
	}
}

func clientStream(client pb.EchoClient) {
	stream, err := client.ClientStreamingEcho(context.Background())
	if err != nil {
		log.Fatalf("send error:%v", err)
	}

	for i := 0; i < 3; i++ {
		err := stream.Send(&pb.EchoRequest{Message: "hello world"})
		if err != nil {
			log.Printf("send error:%v", err)
			continue
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("closeAndRecv() error:%v", err)
	}
	log.Printf("received: %v", resp.GetMessage())
}

func bidirectionalSteram(client pb.EchoClient) {
	var wg sync.WaitGroup

	stream, err := client.BidirectionalStreamEcho(context.Background())
	if err != nil {
		log.Fatalf("client error:%v", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("server closed")
				break
			}
			if err != nil {
				continue
			}
			fmt.Println("receiv data:", req.GetMessage())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			err := stream.Send(&pb.EchoRequest{Message: "hello world"})
			if err != nil {
				log.Printf("send data error:%v", err)
			}
			time.Sleep(time.Second)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Printf("send error:%v \n", err)
			return
		}
	}()

	wg.Wait()
}
