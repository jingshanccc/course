package handler

import (
	"context"
	"course/service/proto/user"
	"course/service/repository"
	"log"
)

type UserHandler struct {
}

func (u *UserHandler) List(ctx context.Context, in *user.PageDto, out *user.PageDto) error {
	users, err := (&repository.UserRepository{}).List(ctx, in)
	if err != nil {
		return err
	}
	out.PageSize = in.PageSize
	out.PageNum = in.PageNum
	out.SortBy = in.SortBy
	out.Order = in.Order
	out.Rows = users

	return nil
}
func (u *UserHandler) Login(ctx context.Context, in *user.User, out *user.User) error {
	log.Println("user login")
	return nil
}

func (u *UserHandler) Logout(ctx context.Context, in *user.User, out *user.User) error {
	log.Println("user logout")
	return nil
}

func (u *UserHandler) Save(ctx context.Context, in *user.User, out *user.User) error {
	log.Println("save user")
	return nil
}
