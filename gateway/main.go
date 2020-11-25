package main

import (
	"course/course-srv/proto/course"
	"course/gateway/route"
	"course/middleware/redis"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	defer redis.RedisClient.Close()
	//initial registry
	r := consul.NewRegistry(registry.Addrs(public.RegistryAddr))

	client := micro.NewService(micro.Name(public.UserClientName))
	//get UserService from registry
	client.Init(micro.Registry(r))
	userService := user.NewUserService(public.UserServiceName, client.Client())

	client = micro.NewService(micro.Name(public.CourseCliName))
	//get CourseService from registry
	client.Init(micro.Registry(r))
	courseService := course.NewCourseService(public.CourseServiceName, client.Client())

	//create web micro service, register in consul, use gin router to handler request
	server := web.NewService(
		web.Name(public.GatewayName),
		web.Address(":4000"),
		web.Handler(route.NewRouter(userService, courseService)),
		web.Registry(r),
	)
	util.PanicIfErr(server.Init(), public.GatewayName)
	util.PanicIfErr(server.Run(), public.GatewayName)
}
