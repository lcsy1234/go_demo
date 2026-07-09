package user

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/internal/model"
	"demo/internal/service"
	"fmt"
)

func (c *ControllerV1) GetTaskDetail(ctx context.Context, req *v1.GetTaskDetailReq) (res *v1.GetTaskDetailRes, err error) {
	out, err := service.User().GetTaskDetail(ctx, model.GetTaskDetailInput{
		UserId: req.UserId,
	})
	fmt.Println(out)
	if err != nil {
		return nil, err
	}
	return &v1.GetTaskDetailRes{
		GetTaskDetailOutput: &out,
	}, nil
}
