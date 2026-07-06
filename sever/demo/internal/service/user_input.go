package service

import "demo/internal/model/entity"

type (
	CreateUserInput struct {
		Name   string
		Age    int
		Height float64
	}
	CreateUserOutput struct {
		Id uint
	}
	UpdateUserInput struct {
		Id     uint
		Name   string
		Age    int
		Height float64
	}
	GetListInput struct {
		Page int
		Size int
	}
	GetListOutput struct {
		List  []*entity.User
		Total int
	}
	GetAllListOutput struct {
		List []*entity.User
	}
)
