package redis

import (
	"course/public"
	"flag"
	"github.com/gomodule/redigo/redis"
	"time"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     30,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	Pool        *redis.Pool
	redisServer = flag.String("redisServer", public.RedisUrl, "")
)

func init() {
	flag.Parse()
	Pool = newPool(*redisServer)
}

//SetString : 存放string到redis
func SetString(key, value string) {
	Pool.Get().Do("SET", key, value)
}

//GetString : 获取key的value
func GetString(key string) (string, error) {
	return redis.String(Pool.Get().Do("GET", key))
}

//DeleteString : 删除key
func DeleteString(key string) {
	Pool.Get().Do("DEL", key)
}
