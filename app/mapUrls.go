package app

import (
	"test3/hariprathap-hp/system_design/TinyURL/controllers/ping"
	"test3/hariprathap-hp/system_design/tinyURL/controllers/urls"
)

func mapURLs() {
	router.GET("/ping", ping.PingHandler)
	router.POST("/list", urls.List)
	router.POST("/create", urls.Create)
	router.GET("/delete", urls.Delete)
}
