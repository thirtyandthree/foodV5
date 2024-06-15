package middleware

import "github.com/go-redis/redis/v8"

// Middleware 中间件，套入redis
type Middleware struct {
	Client *redis.Client
}
