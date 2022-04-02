package app

import (
	"test3/hariprathap-hp/system_design/tinyURL/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	logger.Info("About to start the application...")
	router.LoadHTMLFiles("templates/index.html")
	router.Run(":8080")
}
