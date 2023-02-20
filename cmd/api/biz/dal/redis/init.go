package redis

import (
	"douyin/pkg/consts"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(
		&redis.Options{
			Addr:     consts.RedisAddr,
			Password: "",
			DB:       0,
		},
	)
	_, err = redisdb.Ping().Result()
	return
}

func Init() {
	if err := initRedis(); err != nil {
		panic(err)
	}
}
