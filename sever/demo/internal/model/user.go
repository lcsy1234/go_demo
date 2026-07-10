package model

// import "demo/sever/demo/internal/model/task"

type User struct {
	Id     uint    `json:"id"     dc:"用户ID"`
	Name   string  `json:"name"   dc:"用户姓名"`
	Age    int     `json:"age"    dc:"用户年龄"`
	Height float64 `json:"height" dc:"用户身高"`
}

type GetTaskDetailInput struct {
	UserId uint `json:"userId" in:"query" v:"required|min:1#用户ID不能为空|用户ID无效"`
}
type GetTaskDetailOutput struct {
	*User
	TaskList []*TaskList `json:"taskList" dc:"任务列表"`
}
