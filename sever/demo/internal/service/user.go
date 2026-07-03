// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"demo/internal/model/entity"
)

type (
	IUser interface {
		CreateUser(ctx context.Context, in CreateUserInput) (out CreateUserOutput, err error)
		GetUser(ctx context.Context, id uint) (user *entity.User, err error)
		UpdateUser(ctx context.Context, in UpdateUserInput) (err error)
		DeleteUser(ctx context.Context, id uint) (err error)
		GetList(ctx context.Context, in GetListInput) (out GetListOutput, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
