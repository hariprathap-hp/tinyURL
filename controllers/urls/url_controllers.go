package urls

import (
	"net/http"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	"test3/hariprathap-hp/system_design/tinyURL/services"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var url tinyurl.Url
	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		zlogger.Error("url_controller: func create(), json binding of user input failed with error ", err)
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	zlogger.Info("url_controller: func create(), creating a new tinyurl for the user " + url.UserID + " and url - " + url.OriginalURL)
	result, createErr := services.UrlServices.CreateURL(url)
	if createErr != nil {
		zlogger.Error("url_controller: func create(), creation of tinyurl failed with error ", errors.NewError(createErr.Error))
		//Handle user creation error
		c.JSON(createErr.Status, createErr)
		return
	}
	zlogger.Info("url_controller: func create(), creation of tinyurl succeeded")
	c.JSON(http.StatusCreated, result)
}

func Delete(c *gin.Context) {
	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		zlogger.Error("url_controller: func delete(), json binding of user input failed with error ", err)
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	delErr := services.UrlServices.DeleteURL(url)
	if delErr != nil {
		zlogger.Error("url_controller: func delete(), creation of tinyurl failed with error ", errors.NewError(delErr.Error))
		c.JSON(delErr.Status, delErr)
		return
	}
	zlogger.Info("url_controller: func create(), deletion of user url succeeded")
	c.String(http.StatusOK, "Url Deleted")
}

func List(c *gin.Context) {
	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		zlogger.Error("url_controller: func list(), json binding of user input failed with error ", err)
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	result, listErr := services.UrlServices.GetURL(url.UserID)
	if listErr != nil {
		zlogger.Error("url_controller: func list(), listing of urls for the users failed with error ", errors.NewError(listErr.Error))
		c.JSON(listErr.Status, listErr)
		return
	}
	zlogger.Info("url_controller: func list(), deletion of user url succeeded")
	c.JSON(http.StatusCreated, result)
}
