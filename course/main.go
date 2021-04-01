package main

import (
	"gitee.com/jingshanccc/course/course/handler"
	"gitee.com/jingshanccc/course/course/proto/course"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
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
	r := etcd.NewRegistry(
		registry.Addrs(config.RegistryAddr))
	t, io, err := public.NewTracer(config.CourseServiceName)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing2.SetGlobalTracer(t)
	service := micro.NewService(
		micro.Registry(r),
		micro.Name(config.CourseServiceName),
		micro.Address(config.CourseServiceAddr),
		micro.WrapHandler(opentracing.NewHandlerWrapper(opentracing2.GlobalTracer())),
	)
	service.Init()
	err = course.RegisterCourseServiceHandler(service.Server(), new(handler.CourseServiceHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
