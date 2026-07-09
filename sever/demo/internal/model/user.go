package model

type User struct {
	Id     uint    `json:"id"     dc:"用户ID"`
	Name   string  `json:"name"   dc:"用户姓名"`
	Age    int     `json:"age"    dc:"用户年龄"`
	Height float64 `json:"height" dc:"用户身高"`
}

type TaskList struct {
	Id       uint   `json:"id"       dc:"任务ID"`
	UserId   uint   `json:"userId"   dc:"用户ID"`
	Title    string `json:"title"    dc:"任务标题"`
	Urgency  string `json:"urgency"  dc:"紧急程度"`
	Deadline string `json:"deadline" dc:"截止时间"`
	Status   int    `json:"status"   dc:"任务状态 0:未完成 1:已完成"`
}

type GetTaskDetailInput struct {
	UserId uint `json:"userId" in:"query" v:"required|min:1#用户ID不能为空|用户ID无效"`
}

type GetTaskDetailOutput struct {
	*User
	TaskList []*TaskList `json:"taskList" dc:"任务列表"`
}
