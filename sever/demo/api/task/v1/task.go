package v1

import (
	"demo/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/task/list" method:"get" tags:"Task" summary:"任务分页列表"`
	UserId uint `json:"userId" in:"query" v:"omitempty|min:1#用户ID无效"`
	Page   int  `json:"page"   in:"query" d:"1"  v:"min:1#页码最小为1"`
	Size   int  `json:"size"   in:"query" d:"10" v:"between:1,100#每页数量范围为1-100"`
}

type GetListRes struct {
	List  []*model.TaskList `json:"list"  dc:"任务列表"`
	Total int               `json:"total" dc:"总数"`
}
