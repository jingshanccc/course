# Micah在线视频课程系统

![license](https://gitee.com/jingshanccc/image/raw/master/image/license.svg)![Go version](https://img.shields.io/github/go-mod/go-version/jingshanccc/course)[![stars](https://img.shields.io/github/stars/jingshanccc/course?style=flat&color=red)![forks](https://img.shields.io/github/forks/jingshanccc/course?label=Fork&color=yellow)](https://github.com/jingshanccc/course)![commit](https://img.shields.io/github/last-commit/jingshanccc/course)



## 项目简介

一个基于 Go 1.14、Go-Micro v2、Gorm 1.20.7、Vue、ElementUI的微服务化的前后端分离的在线视频课程系统

## 目录说明

```bash
course/
├─api_test                  接口测试文件
├─course                    课程微服务
|   ├─proto                 proto文件 定义RPC接口和消息
|   ├─handler		    课程微服务RPC接口实现
|   ├─dao		    课程微服务数据库接口实现
├─doc			    项目相关文档
├─file			    文件微服务
|  ├─proto		    proto文件
|  ├─handler		    文件微服务RPC接口实现
|  ├─dao		    文件微服务数据库接口实现
├─gateway		    网关微服务
|    ├─route		    网关HTTP接口
|    ├─middleware           网关中间件 鉴权、拦截等
|    ├─handler		    网关HTTP接口实现
|    |    ├─user	    用户微服务HTTP接口实现
|    |    ├─file	    文件微服务HTTP接口实现
|    |    ├─course	    课程微服务HTTP接口实现
├─public		    项目公共依赖
|   ├─util		    工具类 包含json转换、加解密、uuid等
|   ├─proto		    公共proto文件
|   ├─middleware	    中间件
|   |     ├─redis	    redis
|   ├─config		    公共配置
├─user			    用户微服务
|  ├─proto		    proto文件
|  ├─handler		    用户微服务RPC接口实现
|  ├─dao		    用户微服务数据库接口实现

```
