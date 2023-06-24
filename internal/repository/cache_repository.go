package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
)

type cacheRepo struct {
	redisClient *redis.Client
}

func NewCacheRepository(client *redis.Client) model.CacheRepository {
	return &cacheRepo{redisClient: client}
}

func (c *cacheRepo) Get(ctx context.Context, key string) (string, error) {
	val, err := c.redisClient.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (c *cacheRepo) Set(ctx context.Context, key, val string) error {
	return c.redisClient.Set(ctx, key, val, 0).Err()
}

func (c *cacheRepo) Delete(ctx context.Context, keys ...string) error {
	return c.redisClient.Del(ctx, keys...).Err()
}

func (c *cacheRepo) HashGet(ctx context.Context, hash, key string) (string, error) {
	val, err := c.redisClient.HGet(ctx, hash, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (c *cacheRepo) HashSet(ctx context.Context, hash, key, val string) error {
	return c.redisClient.HSet(ctx, hash, key, val).Err()
}
