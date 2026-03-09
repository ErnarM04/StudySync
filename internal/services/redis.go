package services

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// RedisClient is the process-wide Redis connection used for caching hot reads and shared transient state.
var RedisClient *redis.Client

// Ctx is the default context passed to Redis commands from this package.
var Ctx = context.Background()

// InitRedis dials localhost:6379 (default DB) and should run during application startup before cache use.
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
