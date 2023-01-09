package main

import (
	"context"
	ecpb "demo002/proto/echo"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50053", "the address to connect to")

func logger(format string, a ...interface{}) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}

func unaryInterceptor(cxt context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(cxt, method, req, reply, cc, opts...)
	end := time.Now()
	logger("RPC:%s, req:%+v, start time:%s, end time:%s, err:%v", method, req, start.Format(time.RFC3339), end.Format(time.RFC3339), err)

	return err
}

type wrappedStream struct {
	grpc.ClientStream
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	logger("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptor), grpc.WithStreamInterceptor(streamInterceptor), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	client := ecpb.NewEchoClient(conn)
	callUnaryEcho(client, "hello world")
	callBidiStreamEcho(client)
}

func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho()=_, %v", err)
	}
	fmt.Println("UnaryEcho: ", resp.GetMessage())
}

func callBidiStreamEcho(client ecpb.EchoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := client.BidirectionalStreamingEcho(ctx)
	if err != nil {
		return
	}
	for i := 0; i < 3; i++ {
		if err := c.Send(&ecpb.EchoRequest{Message: fmt.Sprintf("request %d", i+1)}); err != nil {
			log.Fatalf("failed to send request due to error:%v", err)
		}
	}
	if err := c.CloseSend(); err != nil {
		log.Printf("close send err:%v", err)
	}

	for {
		resp, err := c.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response dut to error:%v", err)
		}
		fmt.Println("bidistreaming Echo: ", resp.GetMessage())
	}

}
