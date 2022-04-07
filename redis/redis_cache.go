package appcache

import (
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) APPcache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(keys []string) {
	client := cache.getClient()
	client.LPush("urlapp_cache", keys)
}

func (cache *redisCache) Get() string {
	client := cache.getClient()
	key := client.LPop("urlapp_cache")
	return key.Val()
}
