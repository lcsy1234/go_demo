package user

import (
	"context"

	"demo/api/user/v1"
	"demo/internal/service"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	err = service.User().UpdateUser(ctx, service.UpdateUserInput{
		Id:     req.Id,
		Name:   req.Name,
		Age:    req.Age,
		Height: req.Height,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserRes{}, nil
}
