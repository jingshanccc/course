package repository

import (
	"context"
	"course/service/proto/user"
	"course/service/public"
	"log"
)

type UserRepository struct {
	//
	//Id string `json:"id"`
	//
	//LoginName string `json:"login_name"`
	//
	//Name string `json:"name"`
	//
	//Password string `json:"password"`
}

func (u *UserRepository) List(ctx context.Context, dto *user.PageDto) ([]*user.User, error){
	orderby := "desc"
	if dto.Order == 1 {
		orderby = "asc"
	}
	stmt, err := public.DB.Prepare("select * from user order by ? " + orderby + "limit ?,? ")
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