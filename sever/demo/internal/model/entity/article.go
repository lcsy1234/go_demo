// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure for table article.
type Article struct {
	Id        uint        `json:"id"        orm:"id"         description:"文章ID"`         // 文章ID
	UserId    uint        `json:"userId"    orm:"user_id"    description:"作者用户ID"`       // 作者用户ID
	Title     string      `json:"title"     orm:"title"      description:"标题"`           // 标题
	Content   string      `json:"content"   orm:"content"    description:"内容"`           // 内容
	Status    int         `json:"status"    orm:"status"     description:"状态 1:发布 0:草稿"` // 状态 1:发布 0:草稿
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
}
