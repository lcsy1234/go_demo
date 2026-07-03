package test1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

func (s *sTest1) CreateUser(ctx context.Context, in *service.CreateUserInput) (out *service.CreateUserOutput, err error) {
	id, err := dao.User.Ctx(ctx).Data(do.User{
		Name: in.Name,
		Age:  in.Age,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &service.CreateUserOutput{Id: uint(id)}, nil
}

func (s *sTest1) GetUser(ctx context.Context, id uint) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	return user, nil
}

func (s *sTest1) UpdateUser(ctx context.Context, in *service.UpdateUserInput) (err error) {
	result, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(do.User{
		Name: in.Name,
		Age:  in.Age,
	}).Update()
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	return nil
}

func (s *sTest1) DeleteUser(ctx context.Context, id uint) (err error) {
	result, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Delete()
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	return nil
}

func (s *sTest1) GetUserList(ctx context.Context, in *service.GetUserListInput) (out *service.GetUserListOutput, err error) {
	m := dao.User.Ctx(ctx)
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	var list []*entity.User
	err = m.Page(in.Page, in.Size).
		OrderDesc(dao.User.Columns().Id).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	return &service.GetUserListOutput{
		List:  list,
		Total: total,
	}, nil
}
