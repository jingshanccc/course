package main

import (
	"context"
	"course/service/proto/user"
	"course/service/public"
	"github.com/micro/go-micro/v2/client"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	//client := micro.NewService(micro.Name("go.micro.client.user"))
	//client.Init(micro.Registry(consul.NewRegistry(registry.Addrs(public.RegistryAddr))))
	usercli := user.NewUserService(public.UserServiceName, client.DefaultClient)
	list, err := usercli.List(context.Background(), &user.PageDto{
		PageSize: 10,
		PageNum:  1,
		SortBy:   "Id",
		Order:    -1,
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
