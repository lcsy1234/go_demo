// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        orm:"id"         description:"ç”¨æˆ·ID"`     // ç”¨æˆ·ID
	Name      string      `json:"name"      orm:"name"       description:"å§“å"`       // å§“å
	Age       int         `json:"age"       orm:"age"        description:"å¹´é¾„"`       // å¹´é¾„
	Height    float64     `json:"height"    orm:"height"     description:"èº«é«˜(cm)"`   // èº«é«˜(cm)
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"åˆ›å»ºæ—¶é—´"` // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"æ›´æ–°æ—¶é—´"` // æ›´æ–°æ—¶é—´
}
