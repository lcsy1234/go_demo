// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package test1

import (
	"context"

	"demo/api/test1/v1"
)

type ITest1V1 interface {
	Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
