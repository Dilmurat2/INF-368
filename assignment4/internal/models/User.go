package models

import (
	protos "assingment4/api/v1"
)

type User struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

func UserPbMessageToStruct(user *protos.User) *User {
	return &User{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
}
