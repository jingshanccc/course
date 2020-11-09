package main

import (
	"course/service/proto/user"
	"course/service/public"
	"course/service/user-srv/handler"
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
	"log"
)

func main() {
	defer public.DB.Close()
	log.SetFlags(log.Llongfile)
	//r := consul.NewRegistry(
	//	registry.Addrs(public.RegistryAddr))
	service := micro.NewService(
		//micro.Registry(r),
		micro.Name(public.UserServiceName),
		)
	service.Init()
	err := user.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
		return
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
