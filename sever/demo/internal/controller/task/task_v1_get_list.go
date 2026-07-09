package task

import (
	"context"

	v1 "demo/api/task/v1"
	"demo/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	out, err := service.Task().GetList(ctx, service.TaskGetListInput{
		UserId: req.UserId,
		Page:   req.Page,
		Size:   req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetListRes{
		List:  out.List,
		Total: out.Total,
	}, nil
}
