package services

import (
	app_cache "test3/hariprathap-hp/system_design/tinyURL/domain/redis"
)

var (
	KeyService keyServicesInterface = &keyservices{}
)

type keyservices struct{}

type keyServicesInterface interface {
	Get() string
	Set([]string)
	SetKey(string)
}

func (ks *keyservices) Get() string {
	key := app_cache.Rcache.Get()
	return key
}

func (ks *keyservices) Set(keys []string) {
	app_cache.Rcache.Set(keys)
}

func (ks *keyservices) SetKey(key string) {
	app_cache.Rcache.SetKey(key)
}
