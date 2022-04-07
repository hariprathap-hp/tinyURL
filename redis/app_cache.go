package appcache

type APPcache interface {
	Set([]string)
	Get() string
}
