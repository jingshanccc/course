package dao

import (
	"context"
	"course/public"
	"course/user-srv/proto/user"
	"log"
)

type UserDao struct {
}

func (u *UserDao) List(ctx context.Context, dto *user.PageDto) ([]*user.User, error){
	orderby := "desc"
	if dto.Asc {
		orderby = "asc"
	}
	stmt, err := public.DB.Prepare("select * from user order by ? " + orderby + " limit ?,? ")
	if err != nil {
		log.Println("prepare sql failed, err is " + err.Error())
		return nil, err
	}
	rows, err := stmt.Query(dto.SortBy, (dto.PageNum-1)*dto.PageSize, dto.PageSize)
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, err
	}
	var res []*user.User

	for rows.Next() {
		u := &user.User{}
		err = rows.Scan(&u.Id, &u.LoginName, &u.Name, &u.Password)
		CheckErr(err)
		res = append(res, u)
	}
	return res, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
