package user

import (
	"context"
	"demo/internal/service"

	v1 "demo/api/user/v1"
)

func (c *ControllerV1) GetAllList(ctx context.Context, req *v1.GetAllListReq) (res *v1.GetAllListRes, err error) {
	out, err := service.User().GetAllList(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetAllListRes{List: out.List}, nil
}
