package app

import (
	"test3/hariprathap-hp/system_design/TinyURL/controllers/ping"
	"test3/hariprathap-hp/system_design/tinyURL/controllers/urls"
)

func mapURLs() {
	router.GET("/ping", ping.PingHandler)
	router.GET("/index", urls.ListURLs)
	router.POST("/create", urls.CreateURL)
	router.GET("/delete", urls.DeleteURL)
}
