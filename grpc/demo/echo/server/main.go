package main

import (
	"context"
	pb "demo002/echo/proto"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type Echo struct {
	pb.UnimplementedEchoServer
}

func (e *Echo) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", req.GetMessage())
	return &pb.EchoResponse{Message: "unary " + req.GetMessage()}, nil
}

func (e *Echo) ServerStreamingEcho(req *pb.EchoRequest, stream pb.Echo_ServerStreamingEchoServer) error {
	log.Printf("Received: %v", req.GetMessage())
	for i := 0; i < 3; i++ {
		err := stream.Send(&pb.EchoResponse{Message: req.GetMessage()})
		if err != nil {
			log.Fatalf("send error :%+v", err)
			return err
		}
	}
	return nil
}

func (e *Echo) ClientStreamingEcho(stream pb.Echo_ClientStreamingEchoServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("client closed")
			return stream.SendAndClose(&pb.EchoResponse{Message: "ok"})
		}
		if err != nil {
			return err
		}
		log.Printf("received %v", req.GetMessage())
	}
}

func (e *Echo) BidirectionalStreamEcho(stream pb.Echo_BidirectionalStreamEchoServer) error {
	var (
		waitGroup sync.WaitGroup
		msgCh     = make(chan string)
	)
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		for v := range msgCh {
			err := stream.Send(&pb.EchoResponse{Message: v})
			if err != nil {
				fmt.Println("send error:", err)
				continue
			}
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("receive error:%v", err)
			}
			fmt.Printf("receive: %v \n", req.GetMessage())
			msgCh <- req.GetMessage()
		}
		close(msgCh)
	}()

	waitGroup.Wait()

	return nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &Echo{})

	log.Println("serving grpc on 0.0.0.0" + port)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
