package app

import (
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	zlogger.Info("About to start the application...")
	router.LoadHTMLFiles("templates/index.html")
	router.Run(":8080")
}
