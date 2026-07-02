package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	g.Meta `path:"/test" tags:"test" method:"post" summary:"You first test api"`
	Name string `dc:"姓名"`
	Age int `dc:"年龄"`
}
type TestRes struct {
	Msg  string `json:"msg" dc:"返回消息"`
	Name string `json:"name" dc:"姓名"`
	Age  int    `json:"age" dc:"年龄"`
}