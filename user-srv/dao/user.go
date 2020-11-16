package dao

import (
	"context"
	"course/public"
	"course/user-srv/proto/user"
	"crypto/md5"
	"fmt"
	"log"
)

type UserDao struct {
}

//List : get user page
func (u *UserDao) List(ctx context.Context, dto *user.PageDto) ([]*user.User, public.BusinessException) {
	orderby := "desc"
	if dto.Asc {
		orderby = "asc"
	}
	stmt, err := public.DB.PrepareContext(ctx, "select * from user order by ? "+orderby+" limit ?,? ")
	if err != nil {
		log.Println("prepare sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	rows, err := stmt.Query(dto.SortBy, (dto.PageNum-1)*dto.PageSize, dto.PageSize)
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*user.User

	for rows.Next() {
		u := &user.User{}
		err = rows.Scan(&u.Id, &u.LoginName, &u.Name, &u.Password)
		if err != nil {
			log.Println("row scan failed, err is " + err.Error())
			return nil, public.NewBusinessException(public.ROW_SCAN_ERROR)
		}
		res = append(res, u)
	}
	return res, public.NewBusinessException(public.OK)
}

//SelectByLoginName : get user by login name
func (u *UserDao) SelectByLoginName(ctx context.Context, loginName string) (us *user.User, exception public.BusinessException) {
	stmt, err := public.DB.PrepareContext(ctx, "select * from user where login_name=?")
	if err != nil {
		log.Println("prepare sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	row := stmt.QueryRow(loginName)

	us = &user.User{}
	err = row.Scan(&us.Id, &us.LoginName, &us.Name, &us.Password)
	if err != nil { //can not find user by login name
		log.Println("can not find user by login name, err is " + err.Error())
		return nil, public.NoException("")
	}
	return us, public.NoException("")
}

//Login : login
func (u *UserDao) Login(ctx context.Context, dto *user.User) (*user.LoginUserDto, public.BusinessException) {
	usr, err := u.SelectByLoginName(ctx, dto.LoginName)
	if err.Code() != int32(public.OK) {
		log.Println("select by username failed, err is " + err.Error())
		return nil, err
	}
	if usr == nil {
		err = public.NewBusinessException(public.MEMBER_NOT_EXIST)
		log.Println(err.Error() + dto.LoginName)
		return nil, err
	} else {
		if usr.Password == dto.Password {
			res := &user.LoginUserDto{
				Id:        usr.Id,
				LoginName: usr.LoginName,
				Name:      usr.Name,
			}
			res.Token = public.GetUuid()
			setAuth(res)
			return res, public.NoException("")
		} else {
			err = public.NewBusinessException(public.LOGIN_USER_ERROR)
			log.Println(err.Error() + dto.LoginName)
			return nil, err
		}
	}
}

//setAuth : set user's resources (access control)
func setAuth(loginUser *user.LoginUserDto) {

}

//SavePassword : reset password
func (u *UserDao) SavePassword(ctx context.Context, dto *user.User) (string, public.BusinessException) {
	stmt, err := public.DB.PrepareContext(ctx, "update user set password = ? where login_name = ?")
	if err != nil {
		log.Println("prepare sql failed, err is " + err.Error())
		return "", public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	_, err = stmt.Exec(dto.Password, dto.LoginName)
	//result, err := public.DB.ExecContext(ctx, "", dto.Password, dto.LoginName)
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return "", public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	// if password never changed, rows=0, but should not throw exception
	//if rows, err := result.RowsAffected(); rows <=0 || err!=nil{
	//	return "", public.NewBusinessException(public.ROW_SCAN_ERROR)
	//}
	return dto.Password, public.NoException("")
}

//Save : update when dto.id exists, insert otherwise
func (u *UserDao) Save(ctx context.Context, dto *user.User) (*user.User, public.BusinessException) {
	if dto.Id != "" { //update
		stmt, err := public.DB.PrepareContext(ctx, "update user set login_name = ?, name=? where id = ?")
		if err != nil {
			return &user.User{}, public.NewBusinessException(public.PREPARE_SQL_ERROR)
		}
		defer stmt.Close()
		result, err := stmt.Exec(dto.LoginName, dto.Name, dto.Id)
		if err != nil {
			return &user.User{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
		if rows, err := result.RowsAffected(); rows <= 0 || err != nil {
			return &user.User{}, public.NewBusinessException(public.ROW_SCAN_ERROR)
		}
	} else { //insert
		//if login_name exists, throw error
		us, err := u.SelectByLoginName(ctx, dto.LoginName)
		if err.Code() != int32(public.OK) {
			return &user.User{}, err
		}
		if us != nil {
			return &user.User{}, public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
		}
		dto.Id = public.GetShortUuid()
		dto.Password = fmt.Sprintf("%x", md5.Sum([]byte(dto.Password)))
		stmt, err1 := public.DB.PrepareContext(ctx, "insert into user(id, name, login_name, password) values (?, ?, ?, ?)")
		if err1 != nil {
			return &user.User{}, public.NewBusinessException(public.PREPARE_SQL_ERROR)
		}
		defer stmt.Close()
		result, err1 := stmt.Exec(dto.Id, dto.Name, dto.LoginName, dto.Password)
		if err1 != nil {
			return &user.User{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
		if rows, err1 := result.RowsAffected(); rows <= 0 || err1 != nil {
			return &user.User{}, public.NewBusinessException(public.ROW_SCAN_ERROR)
		}
	}
	return dto, public.NoException("")
}
