package test1

import (
	"context"

	"demo/api/test1/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	user, err := service.Test1().GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserRes{User: user}, nil
}
