package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"demo/internal/model/entity"
)

type CreateUserReq struct {
	g.Meta `path:"/user" method:"post" tags:"User" summary:"创建用户"`
	Name   string  `json:"name"   v:"required|length:1,64#姓名不能为空|姓名长度为1-64个字符"`
	Age    int     `json:"age"    v:"required|between:0,150#年龄不能为空|年龄范围为0-150"`
	Height float64 `json:"height" v:"required|between:50,250#身高不能为空|身高范围为50-250cm"`
}

type CreateUserRes struct {
	Id uint `json:"id" dc:"用户ID"`
}

type GetUserReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"获取用户详情"`
	Id     uint `json:"id" in:"path" v:"required|min:1#用户ID不能为空|用户ID无效"`
}

type GetUserRes struct {
	*entity.User
}

type UpdateUserReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"更新用户"`
	Id     uint    `json:"id"     in:"path" v:"required|min:1#用户ID不能为空|用户ID无效"`
	Name   string  `json:"name"   v:"required|length:1,64#姓名不能为空|姓名长度为1-64个字符"`
	Age    int     `json:"age"    v:"required|between:0,150#年龄不能为空|年龄范围为0-150"`
	Height float64 `json:"height" v:"required|between:50,250#身高不能为空|身高范围为50-250cm"`
}

type UpdateUserRes struct{}

type DeleteUserReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"删除用户"`
	Id     uint `json:"id" in:"path" v:"required|min:1#用户ID不能为空|用户ID无效"`
}

type DeleteUserRes struct{}

type GetListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"User" summary:"用户列表"`
	Page   int `json:"page" in:"query" d:"1"  v:"min:1#页码最小为1"`
	Size   int `json:"size" in:"query" d:"10" v:"between:1,100#每页数量范围为1-100"`
}

type GetListRes struct {
	List  []*entity.User `json:"list"  dc:"用户列表"`
	Total int            `json:"total" dc:"总数"`
}
