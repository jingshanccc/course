package redis

import (
	"course/public"
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     public.RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
