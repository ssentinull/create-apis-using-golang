package model

import "context"

type CacheRepository interface {
	Get(ctx context.Context, key string) (reply string, err error)
	Set(ctx context.Context, key, val string) (err error)
	Delete(ctx context.Context, keys ...string) (err error)
	HashGet(ctx context.Context, hash, key string) (reply string, err error)
	HashSet(ctx context.Context, hash, key, val string) (err error)
}
