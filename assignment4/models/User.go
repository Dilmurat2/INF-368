package models

import "assingment4/proto/pb"

type User struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

func UserPbMessageToStruct(user *pb.User) *User {
	return &User{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
}
