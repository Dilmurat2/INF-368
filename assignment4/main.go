package main

import (
	"assingment4/config"
	"assingment4/grpc/server"
	"assingment4/proto/pb"
	"assingment4/repository"
	"assingment4/services"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
	userRepo, err := repository.NewUserRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}
	userService := services.NewUserService(userRepo)

	handler := server.NewHandler(userService)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, handler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
