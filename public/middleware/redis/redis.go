package redis

import (
	"gitee.com/jingshanccc/course/public/config"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Conf.BasicConfig.BasicHost + config.Conf.BasicConfig.RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
