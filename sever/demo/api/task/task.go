package task

import (
	"context"

	"demo/api/task/v1"
)

type ITaskV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
