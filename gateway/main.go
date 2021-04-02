package main

import (
	"gitee.com/jingshanccc/course/course/proto/course"
	"gitee.com/jingshanccc/course/file/proto/file"
	"gitee.com/jingshanccc/course/gateway/route"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/middleware/redis"
	"gitee.com/jingshanccc/course/user/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcd/v2"
)

func main() {
	defer redis.RedisClient.Close()
	//initial registry
	r := etcd.NewRegistry(registry.Addrs(config.Conf.BasicConfig.BasicHost + config.Conf.BasicConfig.RegistryAddr))

	client := micro.NewService()
	//get UserService from registry
	client.Init(micro.Registry(r))
	userService := user.NewUserService(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, client.Client())

	client = micro.NewService()
	//get CourseService from registry
	client.Init(micro.Registry(r))
	courseService := course.NewCourseService(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, client.Client())

	client = micro.NewService()
	//get FileService from registry
	client.Init(micro.Registry(r))
	fileService := file.NewFileService(config.Conf.BasicConfig.BasicName+config.Conf.Services["file"].Name, client.Client())

	//create web micro service, register in consul, use gin router to handler request
	server := web.NewService(
		web.Name(config.Conf.BasicConfig.BasicName+config.Conf.Services["gateway"].Name),
		web.Address(config.Conf.BasicConfig.BasicHost+config.Conf.Services["gateway"].Addr),
		web.Handler(route.NewRouter(userService, courseService, fileService)),
		web.Registry(r),
	)
	server.Init()
	server.Run()
}
