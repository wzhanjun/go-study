package main

import (
	"encoding/json"
	"log"
	"net"
	user "user/user"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50001"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	user.RegisterUserServer(grpcServer, &UserSerivce{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve :%v", err)
	}

}

type UserSerivce struct {
}

func (u *UserSerivce) UserIndex(ctx context.Context, in *user.UserIndexRequest) (*user.UserIndexResponse, error) {
	printLog("UserIndex", in)

	return &user.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*user.UserEntity{
			{Name: "张三", Age: 20},
			{Name: "王五", Age: 25},
		},
	}, nil
}

func (u *UserSerivce) UserView(ctx context.Context, in *user.UserViewRequest) (*user.UserViewResponse, error) {
	printLog("UserPost", in)

	return &user.UserViewResponse{
		Err: 0,
		Msg: "success",
		Data: &user.UserEntity{
			Name: "张三",
			Age:  20,
		},
	}, nil
}

func (u *UserSerivce) UserPost(cxt context.Context, in *user.UserPostRequest) (*user.UserPostResponse, error) {
	printLog("UserPost", in)

	return &user.UserPostResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

func (u *UserSerivce) UserDelete(ctx context.Context, in *user.UserDeleteRequest) (*user.UserDeleteResponse, error) {
	printLog("UserPost", in)

	return &user.UserDeleteResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

func printLog(method string, params interface{}) {

	json, _ := json.Marshal(params)

	log.Printf("receiver user " + method + " request :" + string(json))
}
