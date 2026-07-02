package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// 接口1：Hello 接口
type HelloReq struct {
	g.Meta `path:"/hello" method:"get" summary:"打招呼接口"`
	Name   string `v:"required" dc:"姓名"`
	Age    int    `v:"required" dc:"年龄"`
}
type HelloRes struct {
	Msg string `dc:"返回消息"`
}

// 接口2：UserInfo 用户信息接口
type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"post" summary:"获取用户信息"`
	UserId int `v:"min:1" dc:"用户ID"`
}
type UserInfoRes struct {
	UserId int    `dc:"用户ID"`
	Name   string `dc:"用户名"`
}

// 控制器结构体
type Hello struct{}

// 接口1 处理方法
func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	r := g.RequestFromCtx(ctx)
    r.Response.Writef(
        "Hello %s! Your Age is %d",
        req.Name,
        req.Age,
    )
    return
}

// 接口2 处理方法
func (Hello) UserInfo(ctx context.Context, req *UserInfoReq) (res *UserInfoRes, err error) {
	res = &UserInfoRes{
		UserId: req.UserId,
		Name:   "测试用户" + gconv.String(req.UserId),
	}
	return
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(new(Hello))
	})
	s.SetPort(8081)
	s.Run()
}