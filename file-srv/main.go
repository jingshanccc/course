package main

import (
	"course/config"
	"course/file-srv/handler"
	"course/file-srv/proto/file"
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
		micro.Name(config.FileServiceName),
	)
	service.Init()
	err := file.RegisterFileServiceHandler(service.Server(), new(handler.FileServiceHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
