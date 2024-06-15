package app

import (
	"foodV5/common/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Addr,
		Password: config.C.Redis.Password, // no password set
		DB:       config.C.Redis.Database, // use default DB
	})
}
