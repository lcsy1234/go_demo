package task

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/service"
)

type sTask struct{}

func init() {
	service.RegisterTask(New())
}

func New() *sTask {
	return &sTask{}
}

func (s *sTask) GetList(ctx context.Context, in service.TaskGetListInput) (out service.TaskGetListOutput, err error) {
	page, size := in.Page, in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	m := dao.Task.Ctx(ctx)
	if in.UserId > 0 {
		m = m.Where(dao.Task.Columns().UserId, in.UserId)
	}
	total, err := m.Count()
	if err != nil {
		return out, err
	}
	var list []*model.TaskList
	err = m.Page(page, size).
		OrderDesc(dao.Task.Columns().Id).
		Scan(&list)
	if err != nil {
		return out, err
	}
	if list == nil {
		list = make([]*model.TaskList, 0)
	}
	return service.TaskGetListOutput{
		List:  list,
		Total: total,
	}, nil
}
