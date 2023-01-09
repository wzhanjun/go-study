package main

import (
	"context"
	"demo002/authentication"
	ecpb "demo002/proto/echo"
	"flag"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50053", "the address to connect to")

func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}

func main() {
	flag.Parse()

	// 构建一个 PerRPCCredentials。
	// perRPC := oauth.NewOauthAccess(fetchToken())
	myAuth := authentication.NewMyAuth()

	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(perRPC))
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(myAuth))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := ecpb.NewEchoClient(conn)

	callUnaryEcho(client, "hello world")
}

// fetchToken 获取授权信息
func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
