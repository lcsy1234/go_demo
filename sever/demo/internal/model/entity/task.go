// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id        uint        `json:"id"        orm:"id"         description:"任务ID"`                 // 任务ID
	UserId    uint        `json:"userId"    orm:"user_id"    description:"所属用户ID"`               // 所属用户ID
	Title     string      `json:"title"     orm:"title"      description:"任务标题"`                 // 任务标题
	Urgency   string      `json:"urgency"   orm:"urgency"    description:"紧急程度 low/medium/high"` // 紧急程度 low/medium/high
	Deadline  *gtime.Time `json:"deadline"  orm:"deadline"   description:"截止时间"`                 // 截止时间
	Status    int         `json:"status"    orm:"status"     description:"状态 0:未完成 1:已完成"`       // 状态 0:未完成 1:已完成
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                 // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                 // 更新时间
}
