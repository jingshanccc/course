package redis

import (
	"course/config"
	"github.com/go-redis/redis/v8"
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
