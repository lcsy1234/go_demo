package test1

import (
	"context"

	"demo/api/test1/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	err = service.Test1().UpdateUser(ctx, &service.UpdateUserInput{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserRes{}, nil
}
