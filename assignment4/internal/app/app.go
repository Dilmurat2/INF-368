package app

import (
	protos "assingment4/api/v1"
	"assingment4/internal/config"
	"assingment4/internal/repository"
	"assingment4/internal/server"
	"assingment4/internal/services"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type app struct {
	server *server.Server
}

func AppInit() (*app, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	log.Println(cfg)
	userRepo, err := repository.NewUserRepository(cfg)
	if err != nil {
		return nil, err
	}
	userService := services.NewUserService(userRepo)
	handler := server.NewHandler(userService)
	return &app{server: handler}, nil
}

func (app *app) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	protos.RegisterUserServiceServer(s, app.server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}
