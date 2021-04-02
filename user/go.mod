module gitee.com/jingshanccc/course/user

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	gitee.com/jingshanccc/course/public v0.0.0-20210402072743-466ecc0fa016
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.5.1
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcd/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1 // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.32.0 // indirect
	google.golang.org/protobuf v1.26.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)
