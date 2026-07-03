package service

import (
	"context"

	"demo/internal/model/entity"
)

type (
	CreateUserInput struct {
		Name string
		Age  int
	}
	CreateUserOutput struct {
		Id uint
	}
	UpdateUserInput struct {
		Id   uint
		Name string
		Age  int
	}
	GetUserListInput struct {
		Page int
		Size int
	}
	GetUserListOutput struct {
		List  []*entity.User
		Total int
	}
)

type ITest1 interface {
	CreateUser(ctx context.Context, in *CreateUserInput) (out *CreateUserOutput, err error)
	GetUser(ctx context.Context, id uint) (user *entity.User, err error)
	UpdateUser(ctx context.Context, in *UpdateUserInput) (err error)
	DeleteUser(ctx context.Context, id uint) (err error)
	GetUserList(ctx context.Context, in *GetUserListInput) (out *GetUserListOutput, err error)
}

var localTest1 ITest1

func Test1() ITest1 {
	if localTest1 == nil {
		panic("implement not found for interface ITest1, forgot register?")
	}
	return localTest1
}

func RegisterTest1(i ITest1) {
	localTest1 = i
}
