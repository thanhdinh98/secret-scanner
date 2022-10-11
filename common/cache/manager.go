package cache

import (
	"context"
	"sync"
	"time"
)

var (
	cacher *CacheRedisHandler = nil
	lock                      = &sync.Mutex{}
)

type (
	Options struct {
		Timeout time.Duration
	}

	Handler interface {
		Exists(ctx context.Context, key string) (bool, error)
		Get(ctx context.Context, key string, result interface{}) error
		Set(ctx context.Context, key string, value interface{}, timeout time.Duration) error
		SetEx(ctx context.Context, key string, value interface{}, options Options) error
		Delete(ctx context.Context, key string) error
	}

	Encoder interface {
		Dump(fromObj interface{}, toObj interface{}) error
		Encode(v interface{}) ([]byte, error)
		Decode(data []byte, v interface{}) error
	}
)

func getCacher() Handler {
	if cacher == nil {
		defer lock.Unlock()
		lock.Lock()
		if cacher == nil {
			cacher = NewRedisHandler()
		}
	}

	return cacher
}

func Get() Handler {
	return getCacher()
}
