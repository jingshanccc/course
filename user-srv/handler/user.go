package handler

import (
	"context"
	"course/user-srv/dao"
	"course/user-srv/proto/user"
)

type UserHandler struct {
}

func (u *UserHandler) List(ctx context.Context, in *user.PageDto, out *user.PageDto) error {
	users, err := (&dao.UserDao{}).List(ctx, in)
	if err != nil {
		return err
	}
	out.PageSize = in.PageSize
	out.PageNum = in.PageNum
	out.SortBy = in.SortBy
	out.Asc = in.Asc
	out.Users = users

	return nil
}