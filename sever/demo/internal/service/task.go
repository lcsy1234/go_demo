package service

import "context"

type (
	ITask interface {
		GetList(ctx context.Context, in TaskGetListInput) (out TaskGetListOutput, err error)
	}
)

var (
	localTask ITask
)

func Task() ITask {
	if localTask == nil {
		panic("implement not found for interface ITask, forgot register?")
	}
	return localTask
}

func RegisterTask(i ITask) {
	localTask = i
}
