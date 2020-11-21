module course

go 1.14

//解决 etcd clientv3 多个类undefined等问题
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
	github.com/google/uuid v1.1.1
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.25.0
)
