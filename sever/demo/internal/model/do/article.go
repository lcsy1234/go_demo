// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta    `orm:"table:article, do:true"`
	Id        any         // 文章ID
	UserId    any         // 作者用户ID
	Title     any         // 标题
	Content   any         // 内容
	Status    any         // 状态 1:发布 0:草稿
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
