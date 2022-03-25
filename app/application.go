package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	router.LoadHTMLFiles("templates/index.html")
	router.Run(":8080")
}
