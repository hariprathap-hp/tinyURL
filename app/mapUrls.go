package app

import (
	"test3/hariprathap-hp/system_design/TinyURL/controllers/ping"
	"test3/hariprathap-hp/system_design/tinyURL/controllers/urls"
)

func mapURLs() {
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", ping.PingHandler)
	router.GET("/index", urls.Index)
	router.GET("/list", urls.ListURLs)
	router.POST("/create", urls.Create)
	router.GET("/delete", urls.Delete)
	router.GET("/redirect", urls.RedirectURL)
}
