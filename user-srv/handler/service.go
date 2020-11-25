package handler

import "course/user-srv/dao"

/*
 when service handler return not nil, the default err code will be 500.
 it leads to the gateway response internal server error.
 so when service handler throws error, we should generate a *errors.Error(go-micro) with a code not 500,
 so that we can let the gateway response in our way.
*/

var (
	userDao     = &dao.UserDao{}
	resourceDao = &dao.ResourceDao{}
	roleDao     = &dao.RoleDao{}
)

type UserServiceHandler struct {
}
