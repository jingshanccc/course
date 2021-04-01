package main

import (
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/user/handler"
	"gitee.com/jingshanccc/course/user/proto/user"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	_ "io"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	//r := consul.NewRegistry(
	//	registry.Addrs(config.RegistryAddr))
	r := etcd.NewRegistry(
		registry.Addrs(config.RegistryAddr))
	t, io, err := public.NewTracer(config.UserServiceName)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing2.SetGlobalTracer(t)
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
