package user

import (
	"context"

	"demo/api/user/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetAllList(ctx context.Context, req *v1.GetAllListReq) (res *v1.GetAllListRes, err error) {
	out, err := service.User().GetAllList(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetAllListRes{List: out.List}, nil
}
