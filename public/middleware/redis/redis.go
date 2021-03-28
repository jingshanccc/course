package redis

import (
	"github.com/go-redis/redis/v8"
	"public/config"
)

var (
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
