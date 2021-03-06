package urls

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	services "test3/hariprathap-hp/system_design/tinyURL/services/url_services"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func Create(c *gin.Context) {
	var url = tinyurl.Url{
		OriginalURL: c.Request.FormValue("long_url"),
		UserID:      c.Request.FormValue("user_id"),
	}

	result, createErr := services.UrlServices.CreateURL(url)
	if createErr != nil {
		zlogger.Error("url_controller: func create(), creation of tinyurl failed with error ", errors.NewError(createErr.Error))
		//Handle user creation error
		c.HTML(http.StatusInternalServerError, "gotocreate.html", createErr)
		return
	}
	zlogger.Info("url_controller: func create(), creation of tinyurl succeeded")
	r, _ := json.MarshalIndent(result, "", "    ")
	c.HTML(http.StatusCreated, "gotocreate.html", string(r))
}

func Delete(c *gin.Context) {
	fmt.Println("going to delete the url")
	var url = tinyurl.Url{
		TinyURL: c.Request.URL.Query().Get("url"),
	}
	fmt.Println("url to delete is -- ", url)

	delErr := services.UrlServices.DeleteURL(url)
	if delErr != nil {
		zlogger.Error("url_controller: func delete(), creation of tinyurl failed with error ", errors.NewError(delErr.Error))
		c.JSON(delErr.Status, delErr)
		return
	}
	zlogger.Info("url_controller: func delete(), deletion of user url succeeded")
	c.JSON(http.StatusOK, "urldeleted")
}

func ListURLs(c *gin.Context) {
	var url = tinyurl.Url{
		UserID: c.Request.FormValue("email"),
	}
	result, listErr := services.UrlServices.ListURL(url.UserID)
	if listErr != nil {
		zlogger.Error("url_controller: func list(), listing of urls for the user failed with error : ", errors.NewError(listErr.Error))
		c.JSON(listErr.Status, listErr)
		return
	}
	zlogger.Info("url_controller: func list(), successfully listing all the urls for the user")
	c.HTML(http.StatusOK, "url_list.html", result)
}

func RedirectURL(c *gin.Context) {
	var url = tinyurl.Url{
		TinyURL: c.Request.URL.Query().Get("url"),
	}
	res, redirectErr := services.UrlServices.Redirect(url)
	if redirectErr != nil {
		zlogger.Error("url_controller: func list(), listing of urls for the user failed with error : ", errors.NewError(redirectErr.Error))
		c.JSON(redirectErr.Status, redirectErr)
		return
	}
	c.Redirect(http.StatusMovedPermanently, *res)
}
