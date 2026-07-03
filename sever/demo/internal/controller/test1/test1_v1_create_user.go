package test1

import (
	"context"

	"demo/api/test1/v1"
	"demo/internal/service"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	out, err := service.Test1().CreateUser(ctx, &service.CreateUserInput{
		Name: req.Name,
		Age:  req.Age,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserRes{Id: out.Id}, nil
}
