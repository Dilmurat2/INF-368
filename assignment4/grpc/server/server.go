package server

import (
	"assingment4/models"
	"assingment4/proto/pb"
	"assingment4/services"
	"context"
)

type Server struct {
	userService services.UserService
}

func NewHandler(userService services.UserService) *Server {
	return &Server{userService: userService}
}

func (h Server) CreateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	user := models.UserPbMessageToStruct(req)
	id, err := h.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{UserId: id}, nil
}

func (h Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, err := h.userService.GetUserById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h Server) ListUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UserList, error) {
	users, err := h.userService.GetUsersList()
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, user := range *users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &pb.UserList{Users: pbUsers}, nil
}
