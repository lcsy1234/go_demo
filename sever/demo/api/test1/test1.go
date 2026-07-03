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
	CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
