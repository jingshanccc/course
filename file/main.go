package main

import (
	"gitee.com/jingshanccc/course/file/handler"
	"gitee.com/jingshanccc/course/file/proto/file"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcd/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	_ "io"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	r := etcd.NewRegistry(
		registry.Addrs(config.Conf.BasicConfig.BasicHost + config.Conf.BasicConfig.RegistryAddr))
	t, io, err := public.NewTracer(config.Conf.BasicConfig.BasicName + config.Conf.Services["file"].Name)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing2.SetGlobalTracer(t)
	service := micro.NewService(
		micro.Registry(r),
		micro.Name(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name),
		micro.Address(config.Conf.BasicConfig.BasicHost+config.Conf.Services["file"].Addr),
		micro.WrapHandler(opentracing.NewHandlerWrapper(opentracing2.GlobalTracer())),
	)
	service.Init()
	err = file.RegisterFileServiceHandler(service.Server(), new(handler.FileServiceHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
