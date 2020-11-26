# Micah在线视频课程系统

![license](https://gitee.com/jingshanccc/image/raw/master/image/license.svg)![Go version](https://img.shields.io/github/go-mod/go-version/jingshanccc/course)[![stars](https://img.shields.io/github/stars/jingshanccc/course?style=flat&color=red)![forks](https://img.shields.io/github/forks/jingshanccc/course?label=Fork&color=yellow)](https://github.com/jingshanccc/course)![commit](https://img.shields.io/github/last-commit/jingshanccc/course)



## 项目简介

一个基于 Go 1.14、Go-Micro v2、Gorm 1.20.7、Vue、ElementUI的微服务化的前后端分离的在线视频课程系统

## 目录说明

```bash
course/
├─api_test/			······接口测试文件
├─course-srv/			······课程服务
│  ├─dao/			······数据库查询
│  ├─handler/			······课程服务RPC接口实现
│  └─proto/			······课程服务ProtoBuf文件夹
│      ├─course/		······课程服务RPC接口声明
│      └─dto/			......课程服务使用的DTO
├─doc/				......项目相关文档
├─gateway/			......网关服务
│  ├─handler/			......网关HTTP接口实现
│  │  ├─course/			......课程服务HTTP接口实现
│  │  └─user/			......用户服务HTTP接口实现
│  └─route/			......网关路由
├─middleware/			......中间件
│  └─redis/			......Redis缓存
├─proto/			......项目ProtoBuf文件夹
│  └─basic/			......项目公用的DTO
├─public/			......公共文件夹
│  └─util/			......工具类
└─user-srv/			......用户服务
    ├─dao/			......数据库查询
    ├─handler/			......RPC接口实现
    └─proto/			......用户服务ProtoBuf文件夹
        ├─dto/			......用户服务使用的DTO
        └─user/			......用户服务RPC接口声明
```
