package urls

import (
	"fmt"
	"net/http"
	"strconv"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	"test3/hariprathap-hp/system_design/tinyURL/services"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"

	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) int64 {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		fmt.Println("error getting user-id")
		return -1
	}
	return user_id
}

func CreateURL(c *gin.Context) {
	var url tinyurl.Url

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

func ListURLs(c *gin.Context) {

	var url tinyurl.Url

	//we can use shouldBindJson instead of json.Marshal
	if err := c.ShouldBindJSON(&url); err != nil {
		restError := errors.NewBadRequestError("Bad JSON Request")
		c.JSON(restError.Status, restError)
		return
	}
	result, listErr := services.GetURL(url.UserID)
	if listErr != nil {
		//Handle user creation error
		c.JSON(listErr.Status, listErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
