package main

import (
	"course/config"
	"course/user-srv/handler"
	"course/user-srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	r := consul.NewRegistry(
		registry.Addrs(config.RegistryAddr))
	service := micro.NewService(
		micro.Registry(r),
		micro.Name(config.UserServiceName),
	)
	service.Init()
	err := user.RegisterUserServiceHandler(service.Server(), new(handler.UserServiceHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
		return
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
