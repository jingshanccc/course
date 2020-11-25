package main

import (
	"course/course-srv/handler"
	"course/course-srv/proto/course"
	"course/public"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	r := consul.NewRegistry(
		registry.Addrs(public.RegistryAddr))
	service := micro.NewService(
		micro.Registry(r),
		micro.Name(public.CourseServiceName),
	)
	service.Init()
	err := course.RegisterCourseServiceHandler(service.Server(), new(handler.CourseServiceHandler))
	if err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
