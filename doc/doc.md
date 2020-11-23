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
