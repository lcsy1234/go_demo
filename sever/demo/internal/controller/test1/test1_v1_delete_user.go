package test1

import (
	"context"

	"demo/api/test1/v1"
	"demo/internal/service"
)

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	err = service.Test1().DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserRes{}, nil
}
