package test1

import (
	"context"

	"demo/api/test1/v1"
)

func (c *ControllerV1) Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	res = &v1.TestRes{
		Msg:  "Test OK!",
		Name: req.Name,
		Age:  req.Age,
	}
	return
}
