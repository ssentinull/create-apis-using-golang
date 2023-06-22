package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/ssentinull/create-apis-using-golang/internal/config"
)

var (
	RedisClient *redis.Client
)

func InitializeRedisConn() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisHost(),
		Password: config.RedisPassword(),
		DB:       config.RedisDB(),
	})
}
