package cache

import (
	"context"
	"guardian/config"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheRedisHandler struct {
	client *redis.Client
	prefix string

	encoder Encoder
}

func NewRedisHandler() *CacheRedisHandler {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(config.KeyDBConnection),
		Password: "",
		DB:       0,
	})

	return &CacheRedisHandler{
		client:  redisClient,
		prefix:  "scanner:",
		encoder: &JsonEncoder{},
	}
}

func (rh *CacheRedisHandler) Get(ctx context.Context, key string, result interface{}) error {
	var (
		fullKey = genKey(rh.prefix, key)
		cmd     = rh.client.Get(ctx, fullKey)
	)
	if cmd.Err() == redis.Nil {
		return ErrNotFound
	}
	bytes, err := cmd.Bytes()
	if err != nil {
		return err
	}
	return rh.encoder.Decode(bytes, &result)
}

func (rh *CacheRedisHandler) Exists(ctx context.Context, key string) (bool, error) {
	var (
		fullKey = genKey(rh.prefix, key)
		cmd     = rh.client.Exists(ctx, fullKey)
	)
	result, err := cmd.Result()
	return result > 0, err
}

func (rh *CacheRedisHandler) Set(
	ctx context.Context,
	key string, value interface{}, timeout time.Duration,
) error {
	return rh.SetEx(ctx, key, value, Options{Timeout: timeout})
}

func (rh *CacheRedisHandler) SetEx(
	ctx context.Context,
	key string, value interface{}, options Options,
) error {
	valueBytes, err := rh.encoder.Encode(value)
	if err != nil {
		return err
	}
	fullKey := genKey(rh.prefix, key)
	return rh.client.Set(ctx, fullKey, valueBytes, options.Timeout).Err()
}

func (rh *CacheRedisHandler) Delete(ctx context.Context, key string) error {
	fullKey := genKey(rh.prefix, key)
	return rh.client.Del(ctx, fullKey).Err()
}
