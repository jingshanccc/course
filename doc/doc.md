
# Git
## 修改上次提交信息 commit message
```bash
# 进入修改 (vim)
git commit --amend
# 修改后退出
:wq
#推送到仓库
git push -f
```
## 拷贝git仓库（包含提交日志）
```bash
# 从原始仓库拷贝一份不包含工作区的裸版本库
git clone --bare git@原仓库地址
# 进入项目目录并推送到新的仓库地址
cd xx
git push --mirror git@目标仓库地址
```

# Proto
```bash
# 位于目录: 项目目录的上级目录
# 编译dto下的文件
protoc --proto_path=. --go_out=. course/user-srv/proto/dto/user.proto
# 编译微服务文件
protoc --proto_path=. --go_out=. --micro_out=. course/user-srv/proto/user/user.proto
```

# Docker
## Consul

```bash
# 启动第一个server节点
docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul agent -server -bootstrap-except 2 -ui -bind=0.0.0.0 -client=0.0.0.0
# 获取第一个节点的ip
docker inspect --format '{{ .NetworkSettings.IPAddress }}' consul1
# 开启其他server节点并加入第一个节点
docker run --name consul2 -d -p 8501:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join x.x.x.x
docker run --name consul3 -d -p 8502:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join x.x.x.x
# 查看集群状态
docker exec -it consul1 consul members
```

## Mysql
```
docker run -itd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

## MariaDB

```
docker run --name mariadb -v /var/lib/mysql:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root@123456 -e MYSQL_ROOT_HOST=% -d mariadb --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

## Redis
```
docker run -itd -p 6379:6379 redis
```

## 自动化发布API
- 基础命令
```bash
docker run -it --rm \
  -e ACCESS_TOKEN={访问令牌} \
  -e APIDOC_TEAM={企业域名} \
  -e APIDOC_PROJECT={项目地址名称} \
  -e APIDOC_ID={API 文档资源 ID} \
  -e APIDOC_RELEASE_TYPE=file \
  -v {API 数据文件路径}/data.txt:/opt/data.txt \
  ecoding/apidoc-publisher
```
- 编写dockerfile构建自己的镜像用于发布API
> 主要是在Dockerfile中，定义需要的环境变量，以及API文件路径。
```bash
# 运行 ecoding/apidoc-publisher，从容器中获取其发布API的脚本
docker cp 容器ID:/my_shell/api_doc_release.sh /xx/
# 修改脚本内容，增加下载文件命令（line 52-53）
```
![](https://gitee.com/jingshanccc/image/raw/master/image/image-20201201120717958.png)
```bash
# 编写Dockerfile
vim Dockerfile
FROM ecoding/apidoc-publisher:20200115.1
ADD api_doc_release.sh /my_shell/api_doc_release.sh
ENV FILE_URL=https://gitee.com/jingshanccc/course/raw/main/doc/openapi.yaml
ENV ACCESS_TOKEN=your token
ENV APIDOC_TEAM=your team
ENV APIDOC_PROJECT=your project
ENV APIDOC_ID=1
ENV APIDOC_RELEASE_TYPE=file
# 构建镜像
docker build -t xx/xxx .
```