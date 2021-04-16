module gitee.com/jingshanccc/course/gateway

go 1.14

require (
	gitee.com/jingshanccc/course/course v0.0.0-20210402092213-0343d5cd5dd6
	gitee.com/jingshanccc/course/file v0.0.0-20210402092213-0343d5cd5dd6
	gitee.com/jingshanccc/course/public v0.0.0-20210402090200-e69eac553a8c
	gitee.com/jingshanccc/course/user v0.0.0-20210402092213-0343d5cd5dd6
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-oauth2/oauth2/v4 v4.1.2
	github.com/go-oauth2/redis/v4 v4.1.1
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcd/v2 v2.9.1
)

replace (
	gitee.com/jingshanccc/course/course => ../course
	gitee.com/jingshanccc/course/file => ../file
	gitee.com/jingshanccc/course/public => ../public
	gitee.com/jingshanccc/course/user => ../user
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
