package model

type TaskList struct {
	Id       uint   `json:"id"       dc:"任务ID"`
	UserId   uint   `json:"userId"   dc:"用户ID"`
	Title    string `json:"title"    dc:"任务标题"`
	Urgency  string `json:"urgency"  dc:"紧急程度"`
	Deadline string `json:"deadline" dc:"截止时间"`
	Status   int    `json:"status"   dc:"任务状态 0:未完成 1:已完成"`
}
