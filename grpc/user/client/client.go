package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"user/user"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()

	userClient := user.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	userIndexResponse, err := userClient.UserIndex(ctx, &user.UserIndexRequest{
		Page:     1,
		PageSize: 10,
	})

	if err != nil {
		log.Printf("user index could not greet: %v", err)
	}

	if userIndexResponse.Err == 0 {
		log.Printf("user index success: %s", userIndexResponse.Msg)
		userEntityList := userIndexResponse.Data
		for _, row := range userEntityList {
			fmt.Println(row.Name, row.Age)
		}
	} else {
		log.Printf("user index error: %d", userIndexResponse.Err)
	}

	userViewRespone, err := userClient.UserView(ctx, &user.UserViewRequest{
		Uid: 1,
	})

	if err != nil {
		log.Printf("user view could not greet: %v", err)
	}

	if userViewRespone.Err == 0 {
		log.Printf("user view success: %s", userViewRespone.Msg)
		userEntity := userViewRespone.Data
		fmt.Println(userEntity.Name, userEntity.Age)
	} else {
		log.Printf("user index error: %d", userIndexResponse.Err)
	}

	// UserPost 请求
	userPostReponse, err := userClient.UserPost(ctx, &user.UserPostRequest{Name: "big_cat", Password: "123456", Age: 29})
	if err != nil {
		log.Printf("user post could not greet: %v", err)
	}

	if userPostReponse.Err == 0 {
		log.Printf("user post success: %s", userPostReponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostReponse.Err)
	}

	// UserDelete 请求
	userDeleteReponse, err := userClient.UserDelete(ctx, &user.UserDeleteRequest{Uid: 1})
	if err != nil {
		log.Printf("user delete could not greet: %v", err)
	}

	if userDeleteReponse.Err == 0 {
		log.Printf("user delete success: %s", userDeleteReponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteReponse.Err)
	}

}
