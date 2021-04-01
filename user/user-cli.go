package main

import (
	"encoding/base64"
	"fmt"
	"gitee.com/jingshanccc/course/public/util"
)

func main() {
	//util.GetRsaKey()
	util.Test()
	txt := "inzf6bcVY347/KnE0U0acNrI2sPI9cYSlWpn1tJfyR+Le/n2TMAnkDP9lDfSXkB8WW698DVJqq+Gt66Bzh13gDXnwlbKBT2GlVsb1I3UYvc5q9AJa6w8cg3tb4qZStbuTKtcdQ/wAhxxSnZ4RJd3Phm2fvIorhwDHi7X3BW21QA="
	bytes, _ := base64.StdEncoding.DecodeString(txt)
	fmt.Println(util.RsaDecrypt(bytes))
}

//func main() {
//	log.SetFlags(log.Llongfile)
//	client := micro.NewService(micro.Name("go.micro.client.user"))
//	client.Init(micro.Registry(consul.NewRegistry(registry.Addrs(config.RegistryAddr))))
//	usercli := user.NewUserService(config.UserServiceName, client.Client())
//	list, err := usercli.List(context.Background(), &user.PageDto{
//		PageSize: 10,
//		PageNum:  1,
//		SortBy:   "Id",
//		Asc:      true,
//	})
//us, err := usercli.Login(context.Background(), &user.User{
//	Id:        "00000001",
//	LoginName: "admin",
//	Name:      "admin",
//	Password:  "123456",
//})
//	if err != nil {
//		log.Fatalln(errors.WithMessage(err, "find user list error"))
//		return
//	}
//	log.Println(list)
//}
