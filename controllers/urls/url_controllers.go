package urls

import (
	"net/http"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	"test3/hariprathap-hp/system_design/tinyURL/services"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}

	result, createErr := services.UrlServices.CreateURL(url)
	if createErr != nil {
		//Handle user creation error
		c.JSON(createErr.Status, createErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Delete(c *gin.Context) {
	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	delErr := services.UrlServices.DeleteURL(url)
	if delErr != nil {
		c.JSON(delErr.Status, delErr)
	}
	c.String(http.StatusOK, "Url Deleted")
}

func List(c *gin.Context) {

	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	result, listErr := services.UrlServices.GetURL(url.UserID)
	if listErr != nil {
		//Handle user creation error
		c.JSON(listErr.Status, listErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
