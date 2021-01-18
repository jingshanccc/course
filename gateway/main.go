package main

import (
	"course/config"
	"course/course-srv/proto/course"
	"course/file-srv/proto/file"
	"course/gateway/route"
	"course/middleware/redis"
	"course/user-srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	defer redis.RedisClient.Close()
	//initial registry
	r := consul.NewRegistry(registry.Addrs(config.RegistryAddr))

	client := micro.NewService(micro.Name(config.UserClientName))
	//get UserService from registry
	client.Init(micro.Registry(r))
	userService := user.NewUserService(config.UserServiceName, client.Client())

	client = micro.NewService(micro.Name(config.CourseCliName))
	//get CourseService from registry
	client.Init(micro.Registry(r))
	courseService := course.NewCourseService(config.CourseServiceName, client.Client())

	client = micro.NewService(micro.Name(config.FileClientName))
	//get CourseService from registry
	client.Init(micro.Registry(r))
	fileService := file.NewFileService(config.FileServiceName, client.Client())

	//create web micro service, register in consul, use gin router to handler request
	server := web.NewService(
		web.Name(config.GatewayName),
		web.Address(":4000"),
		web.Handler(route.NewRouter(userService, courseService, fileService)),
		web.Registry(r),
	)
	server.Init()
	server.Run()
}
