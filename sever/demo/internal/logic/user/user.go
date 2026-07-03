package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) CreateUser(ctx context.Context, in service.CreateUserInput) (out service.CreateUserOutput, err error) {
	id, err := dao.User.Ctx(ctx).Data(do.User{
		Name:   in.Name,
		Age:    in.Age,
		Height: in.Height,
	}).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return service.CreateUserOutput{Id: uint(id)}, nil
}

func (s *sUser) GetUser(ctx context.Context, id uint) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	return user, nil
}

func (s *sUser) UpdateUser(ctx context.Context, in service.UpdateUserInput) (err error) {
	result, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(do.User{
		Name:   in.Name,
		Age:    in.Age,
		Height: in.Height,
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

func (s *sUser) DeleteUser(ctx context.Context, id uint) (err error) {
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

func (s *sUser) GetList(ctx context.Context, in service.GetListInput) (out service.GetListOutput, err error) {
	page, size := in.Page, in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	m := dao.User.Ctx(ctx)
	total, err := m.Count()
	if err != nil {
		return out, err
	}
	var list []*entity.User
	err = m.Page(page, size).
		OrderDesc(dao.User.Columns().Id).
		Scan(&list)
	if err != nil {
		return out, err
	}
	return service.GetListOutput{
		List:  list,
		Total: total,
	}, nil
}
