package urls

import (
	"net/http"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	"test3/hariprathap-hp/system_design/tinyURL/services"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"

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
		return
	}
	c.String(http.StatusOK, "Url Deleted")
}

func List(c *gin.Context) {
	zlogger.Info("API call to display the list of all the urls associated with the user")
	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		zlogger.Error("listing the urls failed due to bad json request sent by user with error - ", err)
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	result, listErr := services.UrlServices.GetURL(url.UserID)
	if listErr != nil {
		zlogger.Error("listing the urls failed due to bad json request sent by user with error - ", errors.NewError(listErr.Error))
		c.JSON(listErr.Status, listErr)
		return
	}
	zlogger.Info("listing the urls associated with the user")
	c.JSON(http.StatusCreated, result)
}
