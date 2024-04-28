package server

import (
	"assingment4/api/v1"
	"assingment4/internal/models"
	"assingment4/internal/services"
	"context"
)

type Server struct {
	v1.UnimplementedUserServiceServer
	userService services.UserService
}

func NewHandler(userService services.UserService) *Server {
	return &Server{userService: userService}
}

func (h Server) CreateUser(ctx context.Context, req *v1.User) (*v1.CreateUserResponse, error) {
	user := models.UserPbMessageToStruct(req)
	id, err := h.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserResponse{UserId: id}, nil
}

func (h Server) GetUser(ctx context.Context, req *v1.GetUserByIdRequest) (*v1.User, error) {
	user, err := h.userService.GetUserById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h Server) ListUsers(ctx context.Context, req *v1.EmptyRequest) (*v1.UserList, error) {
	users, err := h.userService.GetUsersList()
	if err != nil {
		return nil, err
	}
	var protoUsers []*v1.User
	for _, user := range *users {
		protoUsers = append(protoUsers, &v1.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &v1.UserList{Users: protoUsers}, nil
}
