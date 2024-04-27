package server

import (
	protos "assingment4/internal/api/v1"
	"assingment4/internal/models"
	"assingment4/internal/services"
	"context"
)

type Server struct {
	protos.UnimplementedUserServiceServer
	userService services.UserService
}

func NewHandler(userService services.UserService) *Server {
	return &Server{userService: userService}
}

func (h Server) CreateUser(ctx context.Context, req *protos.User) (*protos.CreateUserResponse, error) {
	user := models.UserPbMessageToStruct(req)
	id, err := h.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &protos.CreateUserResponse{UserId: id}, nil
}

func (h Server) GetUser(ctx context.Context, req *protos.GetUserByIdRequest) (*protos.User, error) {
	user, err := h.userService.GetUserById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &protos.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h Server) ListUsers(ctx context.Context, req *protos.EmptyRequest) (*protos.UserList, error) {
	users, err := h.userService.GetUsersList()
	if err != nil {
		return nil, err
	}
	var protoUsers []*protos.User
	for _, user := range *users {
		protoUsers = append(protoUsers, &protos.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &protos.UserList{Users: protoUsers}, nil
}
