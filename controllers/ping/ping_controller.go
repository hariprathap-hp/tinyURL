package ping

import (
	"net/http"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	zlogger.Info("testing ping on the tinyurl server")
	c.String(http.StatusOK, "pong")
}
