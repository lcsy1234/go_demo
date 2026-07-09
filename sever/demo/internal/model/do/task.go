// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta    `orm:"table:task, do:true"`
	Id        any         // 任务ID
	UserId    any         // 所属用户ID
	Title     any         // 任务标题
	Urgency   any         // 紧急程度 low/medium/high
	Deadline  *gtime.Time // 截止时间
	Status    any         // 状态 0:未完成 1:已完成
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
