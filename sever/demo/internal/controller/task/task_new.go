package task

import (
	"demo/api/task"
)

type ControllerV1 struct{}

func NewV1() task.ITaskV1 {
	return &ControllerV1{}
}
