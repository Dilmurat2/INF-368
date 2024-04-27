package services

import (
	"assingment4/internal/models"
	"assingment4/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) (int32, error)
	GetUserById(id int32) (*models.User, error)
	GetUsersList() (*[]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u userService) CreateUser(user *models.User) (int32, error) {
	return u.repo.CreateUser(user)
}

func (u userService) GetUserById(id int32) (*models.User, error) {
	return u.repo.GetUserById(id)
}

func (u userService) GetUsersList() (*[]models.User, error) {
	return u.repo.GetUsersList()
}
