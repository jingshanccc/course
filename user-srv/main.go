package main

import (
	"course/config"
	"course/public"
	"course/user-srv/handler"
	"course/user-srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	r := consul.NewRegistry(
		registry.Addrs(config.RegistryAddr))
	tracer, io, err := public.NewTracer(config.UserServiceName)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	defer io.Close()
	opentracing2.SetGlobalTracer(tracer)
	service := micro.NewService(
		micro.Registry(r),
		micro.Name(config.UserServiceName),
		micro.Address(config.UserServiceAddr),
		micro.WrapHandler(opentracing.NewHandlerWrapper(opentracing2.GlobalTracer())),
	)
	service.Init()
	err = user.RegisterUserServiceHandler(service.Server(), new(handler.UserServiceHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
