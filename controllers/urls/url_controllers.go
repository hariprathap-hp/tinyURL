package urls

import (
	"net/http"
	"test3/hariprathap-hp/minderaWeatherService/utils/errors"
	"test3/hariprathap-hp/system_design/TinyURL/domain/urls"
	"test3/hariprathap-hp/system_design/TinyURL/services"

	"github.com/gin-gonic/gin"
)

func CreateURL(c *gin.Context) {
	var url urls.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}

	result, createErr := services.CreateURL(url)
	if createErr != nil {
		//Handle user creation error
		c.JSON(createErr.Status, createErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func DeleteURL(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me")
}

func IndexURL(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
