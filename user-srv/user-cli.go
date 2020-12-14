package main

import (
	"context"
	"course/config"
	"course/user-srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	client := micro.NewService(micro.Name("go.micro.client.user"))
	client.Init(micro.Registry(consul.NewRegistry(registry.Addrs(config.RegistryAddr))))
	usercli := user.NewUserService(config.UserServiceName, client.Client())
	list, err := usercli.List(context.Background(), &user.PageDto{
		PageSize: 10,
		PageNum:  1,
		SortBy:   "Id",
		Asc:      true,
	})
	//us, err := usercli.Login(context.Background(), &user.User{
	//	Id:        "00000001",
	//	LoginName: "admin",
	//	Name:      "admin",
	//	Password:  "123456",
	//})
	if err != nil {
		log.Fatalln(errors.WithMessage(err, "find user list error"))
		return
	}
	log.Println(list)
}
