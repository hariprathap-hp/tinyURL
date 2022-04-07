package services

import (
	appcache "test3/hariprathap-hp/system_design/tinyURL/redis"
)

var app_cache appcache.APPcache

func init() {
	app_cache = appcache.NewRedisCache("localhost:6379", 1, 10)
}

var (
	KeyService keyServicesInterface = &keyservices{}
)

type keyservices struct{}

type keyServicesInterface interface {
	Get() string
	Set([]string)
}

func (ks *keyservices) Get() string {
	key := app_cache.Get()
	return key
}

func (ks *keyservices) Set(keys []string) {
	app_cache.Set(keys)
}
