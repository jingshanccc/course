// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: course/user-srv/proto/user/user.proto

package user

import (
	basic "course/proto/basic"
	dto "course/user-srv/proto/dto"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	//----------- 用户接口 -------------
	List(ctx context.Context, in *dto.PageDto, opts ...client.CallOption) (*dto.PageDto, error)
	Save(ctx context.Context, in *dto.UserDto, opts ...client.CallOption) (*dto.UserDto, error)
	SaveUserInfo(ctx context.Context, in *dto.UserDto, opts ...client.CallOption) (*dto.UserDto, error)
	Delete(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	SavePassword(ctx context.Context, in *dto.UpdatePass, opts ...client.CallOption) (*basic.String, error)
	Login(ctx context.Context, in *dto.LoginUserDto, opts ...client.CallOption) (*dto.LoginUserDto, error)
	Logout(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.String, error)
	UserInfo(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.UserDto, error)
	SaveEmail(ctx context.Context, in *dto.UpdateEmail, opts ...client.CallOption) (*basic.String, error)
	//---------- 权限管理 ---------------
	//resource
	MenuList(ctx context.Context, in *dto.ResourcePageDto, opts ...client.CallOption) (*dto.ResourcePageDto, error)
	MenuParent(ctx context.Context, in *basic.IntegerList, opts ...client.CallOption) (*dto.ResourceDtoList, error)
	LoadMenus(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.ResourceDtoList, error)
	LoadTree(ctx context.Context, in *basic.Integer, opts ...client.CallOption) (*dto.ResourceDtoList, error)
	MenuChild(ctx context.Context, in *basic.Integer, opts ...client.CallOption) (*basic.IntegerList, error)
	SaveResource(ctx context.Context, in *dto.ResourceDto, opts ...client.CallOption) (*basic.String, error)
	DeleteResource(ctx context.Context, in *basic.IntegerList, opts ...client.CallOption) (*basic.String, error)
	//role
	RoleList(ctx context.Context, in *dto.RolePageDto, opts ...client.CallOption) (*dto.RolePageDto, error)
	AllRole(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.RoleDtoList, error)
	GetRole(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.RoleDto, error)
	RoleLevel(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.Integer, error)
	SaveRole(ctx context.Context, in *dto.RoleDto, opts ...client.CallOption) (*dto.RoleDto, error)
	DeleteRole(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	SaveRoleResource(ctx context.Context, in *dto.RoleDto, opts ...client.CallOption) (*dto.RoleDto, error)
	ListRoleResource(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.IntegerList, error)
	SaveRoleUser(ctx context.Context, in *dto.RoleDto, opts ...client.CallOption) (*dto.RoleDto, error)
	ListRoleUser(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.StringList, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) List(ctx context.Context, in *dto.PageDto, opts ...client.CallOption) (*dto.PageDto, error) {
	req := c.c.NewRequest(c.name, "UserService.List", in)
	out := new(dto.PageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Save(ctx context.Context, in *dto.UserDto, opts ...client.CallOption) (*dto.UserDto, error) {
	req := c.c.NewRequest(c.name, "UserService.Save", in)
	out := new(dto.UserDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SaveUserInfo(ctx context.Context, in *dto.UserDto, opts ...client.CallOption) (*dto.UserDto, error) {
	req := c.c.NewRequest(c.name, "UserService.SaveUserInfo", in)
	out := new(dto.UserDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SavePassword(ctx context.Context, in *dto.UpdatePass, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.SavePassword", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *dto.LoginUserDto, opts ...client.CallOption) (*dto.LoginUserDto, error) {
	req := c.c.NewRequest(c.name, "UserService.Login", in)
	out := new(dto.LoginUserDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Logout(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.Logout", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.UserDto, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(dto.UserDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SaveEmail(ctx context.Context, in *dto.UpdateEmail, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.SaveEmail", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MenuList(ctx context.Context, in *dto.ResourcePageDto, opts ...client.CallOption) (*dto.ResourcePageDto, error) {
	req := c.c.NewRequest(c.name, "UserService.MenuList", in)
	out := new(dto.ResourcePageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MenuParent(ctx context.Context, in *basic.IntegerList, opts ...client.CallOption) (*dto.ResourceDtoList, error) {
	req := c.c.NewRequest(c.name, "UserService.MenuParent", in)
	out := new(dto.ResourceDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) LoadMenus(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.ResourceDtoList, error) {
	req := c.c.NewRequest(c.name, "UserService.LoadMenus", in)
	out := new(dto.ResourceDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) LoadTree(ctx context.Context, in *basic.Integer, opts ...client.CallOption) (*dto.ResourceDtoList, error) {
	req := c.c.NewRequest(c.name, "UserService.LoadTree", in)
	out := new(dto.ResourceDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) MenuChild(ctx context.Context, in *basic.Integer, opts ...client.CallOption) (*basic.IntegerList, error) {
	req := c.c.NewRequest(c.name, "UserService.MenuChild", in)
	out := new(basic.IntegerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SaveResource(ctx context.Context, in *dto.ResourceDto, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.SaveResource", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteResource(ctx context.Context, in *basic.IntegerList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteResource", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RoleList(ctx context.Context, in *dto.RolePageDto, opts ...client.CallOption) (*dto.RolePageDto, error) {
	req := c.c.NewRequest(c.name, "UserService.RoleList", in)
	out := new(dto.RolePageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AllRole(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.RoleDtoList, error) {
	req := c.c.NewRequest(c.name, "UserService.AllRole", in)
	out := new(dto.RoleDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetRole(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.RoleDto, error) {
	req := c.c.NewRequest(c.name, "UserService.GetRole", in)
	out := new(dto.RoleDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RoleLevel(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.Integer, error) {
	req := c.c.NewRequest(c.name, "UserService.RoleLevel", in)
	out := new(basic.Integer)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SaveRole(ctx context.Context, in *dto.RoleDto, opts ...client.CallOption) (*dto.RoleDto, error) {
	req := c.c.NewRequest(c.name, "UserService.SaveRole", in)
	out := new(dto.RoleDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteRole(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteRole", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SaveRoleResource(ctx context.Context, in *dto.RoleDto, opts ...client.CallOption) (*dto.RoleDto, error) {
	req := c.c.NewRequest(c.name, "UserService.SaveRoleResource", in)
	out := new(dto.RoleDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ListRoleResource(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.IntegerList, error) {
	req := c.c.NewRequest(c.name, "UserService.ListRoleResource", in)
	out := new(basic.IntegerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SaveRoleUser(ctx context.Context, in *dto.RoleDto, opts ...client.CallOption) (*dto.RoleDto, error) {
	req := c.c.NewRequest(c.name, "UserService.SaveRoleUser", in)
	out := new(dto.RoleDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ListRoleUser(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.StringList, error) {
	req := c.c.NewRequest(c.name, "UserService.ListRoleUser", in)
	out := new(basic.StringList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//----------- 用户接口 -------------
	List(context.Context, *dto.PageDto, *dto.PageDto) error
	Save(context.Context, *dto.UserDto, *dto.UserDto) error
	SaveUserInfo(context.Context, *dto.UserDto, *dto.UserDto) error
	Delete(context.Context, *basic.StringList, *basic.String) error
	SavePassword(context.Context, *dto.UpdatePass, *basic.String) error
	Login(context.Context, *dto.LoginUserDto, *dto.LoginUserDto) error
	Logout(context.Context, *basic.String, *basic.String) error
	UserInfo(context.Context, *basic.String, *dto.UserDto) error
	SaveEmail(context.Context, *dto.UpdateEmail, *basic.String) error
	//---------- 权限管理 ---------------
	//resource
	MenuList(context.Context, *dto.ResourcePageDto, *dto.ResourcePageDto) error
	MenuParent(context.Context, *basic.IntegerList, *dto.ResourceDtoList) error
	LoadMenus(context.Context, *basic.String, *dto.ResourceDtoList) error
	LoadTree(context.Context, *basic.Integer, *dto.ResourceDtoList) error
	MenuChild(context.Context, *basic.Integer, *basic.IntegerList) error
	SaveResource(context.Context, *dto.ResourceDto, *basic.String) error
	DeleteResource(context.Context, *basic.IntegerList, *basic.String) error
	//role
	RoleList(context.Context, *dto.RolePageDto, *dto.RolePageDto) error
	AllRole(context.Context, *basic.String, *dto.RoleDtoList) error
	GetRole(context.Context, *basic.String, *dto.RoleDto) error
	RoleLevel(context.Context, *basic.String, *basic.Integer) error
	SaveRole(context.Context, *dto.RoleDto, *dto.RoleDto) error
	DeleteRole(context.Context, *basic.StringList, *basic.String) error
	SaveRoleResource(context.Context, *dto.RoleDto, *dto.RoleDto) error
	ListRoleResource(context.Context, *basic.String, *basic.IntegerList) error
	SaveRoleUser(context.Context, *dto.RoleDto, *dto.RoleDto) error
	ListRoleUser(context.Context, *basic.String, *basic.StringList) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		List(ctx context.Context, in *dto.PageDto, out *dto.PageDto) error
		Save(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error
		SaveUserInfo(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error
		Delete(ctx context.Context, in *basic.StringList, out *basic.String) error
		SavePassword(ctx context.Context, in *dto.UpdatePass, out *basic.String) error
		Login(ctx context.Context, in *dto.LoginUserDto, out *dto.LoginUserDto) error
		Logout(ctx context.Context, in *basic.String, out *basic.String) error
		UserInfo(ctx context.Context, in *basic.String, out *dto.UserDto) error
		SaveEmail(ctx context.Context, in *dto.UpdateEmail, out *basic.String) error
		MenuList(ctx context.Context, in *dto.ResourcePageDto, out *dto.ResourcePageDto) error
		MenuParent(ctx context.Context, in *basic.IntegerList, out *dto.ResourceDtoList) error
		LoadMenus(ctx context.Context, in *basic.String, out *dto.ResourceDtoList) error
		LoadTree(ctx context.Context, in *basic.Integer, out *dto.ResourceDtoList) error
		MenuChild(ctx context.Context, in *basic.Integer, out *basic.IntegerList) error
		SaveResource(ctx context.Context, in *dto.ResourceDto, out *basic.String) error
		DeleteResource(ctx context.Context, in *basic.IntegerList, out *basic.String) error
		RoleList(ctx context.Context, in *dto.RolePageDto, out *dto.RolePageDto) error
		AllRole(ctx context.Context, in *basic.String, out *dto.RoleDtoList) error
		GetRole(ctx context.Context, in *basic.String, out *dto.RoleDto) error
		RoleLevel(ctx context.Context, in *basic.String, out *basic.Integer) error
		SaveRole(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error
		DeleteRole(ctx context.Context, in *basic.StringList, out *basic.String) error
		SaveRoleResource(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error
		ListRoleResource(ctx context.Context, in *basic.String, out *basic.IntegerList) error
		SaveRoleUser(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error
		ListRoleUser(ctx context.Context, in *basic.String, out *basic.StringList) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) List(ctx context.Context, in *dto.PageDto, out *dto.PageDto) error {
	return h.UserServiceHandler.List(ctx, in, out)
}

func (h *userServiceHandler) Save(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error {
	return h.UserServiceHandler.Save(ctx, in, out)
}

func (h *userServiceHandler) SaveUserInfo(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error {
	return h.UserServiceHandler.SaveUserInfo(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) SavePassword(ctx context.Context, in *dto.UpdatePass, out *basic.String) error {
	return h.UserServiceHandler.SavePassword(ctx, in, out)
}

func (h *userServiceHandler) Login(ctx context.Context, in *dto.LoginUserDto, out *dto.LoginUserDto) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *userServiceHandler) Logout(ctx context.Context, in *basic.String, out *basic.String) error {
	return h.UserServiceHandler.Logout(ctx, in, out)
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *basic.String, out *dto.UserDto) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}

func (h *userServiceHandler) SaveEmail(ctx context.Context, in *dto.UpdateEmail, out *basic.String) error {
	return h.UserServiceHandler.SaveEmail(ctx, in, out)
}

func (h *userServiceHandler) MenuList(ctx context.Context, in *dto.ResourcePageDto, out *dto.ResourcePageDto) error {
	return h.UserServiceHandler.MenuList(ctx, in, out)
}

func (h *userServiceHandler) MenuParent(ctx context.Context, in *basic.IntegerList, out *dto.ResourceDtoList) error {
	return h.UserServiceHandler.MenuParent(ctx, in, out)
}

func (h *userServiceHandler) LoadMenus(ctx context.Context, in *basic.String, out *dto.ResourceDtoList) error {
	return h.UserServiceHandler.LoadMenus(ctx, in, out)
}

func (h *userServiceHandler) LoadTree(ctx context.Context, in *basic.Integer, out *dto.ResourceDtoList) error {
	return h.UserServiceHandler.LoadTree(ctx, in, out)
}

func (h *userServiceHandler) MenuChild(ctx context.Context, in *basic.Integer, out *basic.IntegerList) error {
	return h.UserServiceHandler.MenuChild(ctx, in, out)
}

func (h *userServiceHandler) SaveResource(ctx context.Context, in *dto.ResourceDto, out *basic.String) error {
	return h.UserServiceHandler.SaveResource(ctx, in, out)
}

func (h *userServiceHandler) DeleteResource(ctx context.Context, in *basic.IntegerList, out *basic.String) error {
	return h.UserServiceHandler.DeleteResource(ctx, in, out)
}

func (h *userServiceHandler) RoleList(ctx context.Context, in *dto.RolePageDto, out *dto.RolePageDto) error {
	return h.UserServiceHandler.RoleList(ctx, in, out)
}

func (h *userServiceHandler) AllRole(ctx context.Context, in *basic.String, out *dto.RoleDtoList) error {
	return h.UserServiceHandler.AllRole(ctx, in, out)
}

func (h *userServiceHandler) GetRole(ctx context.Context, in *basic.String, out *dto.RoleDto) error {
	return h.UserServiceHandler.GetRole(ctx, in, out)
}

func (h *userServiceHandler) RoleLevel(ctx context.Context, in *basic.String, out *basic.Integer) error {
	return h.UserServiceHandler.RoleLevel(ctx, in, out)
}

func (h *userServiceHandler) SaveRole(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	return h.UserServiceHandler.SaveRole(ctx, in, out)
}

func (h *userServiceHandler) DeleteRole(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.UserServiceHandler.DeleteRole(ctx, in, out)
}

func (h *userServiceHandler) SaveRoleResource(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	return h.UserServiceHandler.SaveRoleResource(ctx, in, out)
}

func (h *userServiceHandler) ListRoleResource(ctx context.Context, in *basic.String, out *basic.IntegerList) error {
	return h.UserServiceHandler.ListRoleResource(ctx, in, out)
}

func (h *userServiceHandler) SaveRoleUser(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	return h.UserServiceHandler.SaveRoleUser(ctx, in, out)
}

func (h *userServiceHandler) ListRoleUser(ctx context.Context, in *basic.String, out *basic.StringList) error {
	return h.UserServiceHandler.ListRoleUser(ctx, in, out)
}
