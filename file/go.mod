module gitee.com/jingshanccc/course/file

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	gitee.com/jingshanccc/course/public v0.0.0-20210402090200-e69eac553a8c
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcd/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.26.0
)
