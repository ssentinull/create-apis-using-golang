package db

import "github.com/redis/go-redis/v9"

var (
	RedisClient *redis.Client
)

func InitializeRedisConn() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
