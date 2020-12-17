module course

go 1.14

//解决 etcd clientv3 多个类undefined等问题
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-oauth2/oauth2/v4 v4.1.2
	github.com/go-oauth2/redis/v4 v4.1.1
	github.com/go-redis/redis/v8 v8.0.0-beta.5
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/jordan-wright/email v0.0.0-20200917010138-e1c00e156980
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/mojocn/base64Captcha v1.2.2
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.7
)
