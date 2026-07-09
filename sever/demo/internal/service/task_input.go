package service

import "demo/internal/model"

type (
	TaskGetListInput struct {
		UserId uint
		Page   int
		Size   int
	}
	TaskGetListOutput struct {
		List  []*model.TaskList
		Total int
	}
)
