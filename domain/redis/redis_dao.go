package app_cache

import (
	"test3/hariprathap-hp/system_design/tinyURL/dataResources/redis/keys_db"
	"time"
)

var (
	Rcache appCacheinterface = &redisCache{}
)

type appCacheinterface interface {
	Set([]string)
	Get() string
	SetKey(key string)
}

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) redisCache {
	return redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) Set(keys []string) {
	keys_db.Client.LPush("urlapp_cache", keys)
}

func (cache *redisCache) SetKey(key string) {
	keys_db.Client.LPush("urlapp_cache", key)
}

func (cache *redisCache) Get() string {
	key := keys_db.Client.LPop("urlapp_cache")
	return key.Val()
}
