package main

import (
	"assingment4/api/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	client := v1.NewUserServiceClient(conn)

	// create user
	userId, err := client.CreateUser(context.Background(), &v1.User{
		Name:  "John",
		Email: "johnDokej@gmail.com",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	// get user by id
	user, err := client.GetUser(context.Background(), &v1.GetUserByIdRequest{
		Id: userId.GetUserId(),
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("user: %v", user)
	// get users list
	usersList, err := client.ListUsers(context.Background(), &v1.EmptyRequest{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("users list:", usersList.GetUsers())
}
