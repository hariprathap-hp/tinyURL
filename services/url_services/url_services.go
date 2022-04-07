package services

import (
	"encoding/json"
	"fmt"
	"strings"
	"test3/github.com/mercadolibre/golang-restclient/rest"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	services "test3/hariprathap-hp/system_design/tinyURL/services/cache"
	"test3/hariprathap-hp/system_design/utils_repo/dateutils"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"
	"time"
)

var (
	UrlServices   urlServicesInterface = &urlService{}
	kgsRestClient                      = rest.RequestBuilder{
		BaseURL: "http://localhost:8085",
		Timeout: 100 * time.Second,
	}
)

type urlService struct{}

type urlServicesInterface interface {
	CreateURL(tinyurl.Url) (*tinyurl.Url, *errors.RestErr)
	GetURL(string) (tinyurl.Urls, *errors.RestErr)
	DeleteURL(tinyurl.Url) *errors.RestErr
}

func (u *urlService) CreateURL(url tinyurl.Url) (*tinyurl.Url, *errors.RestErr) {
	if validateErr := url.Validate(); validateErr != nil {
		zlogger.Error("url_service: func create(), validation of user input failed with error - ", errors.NewError(validateErr.Error))
		return nil, validateErr
	}
	zlogger.Info("url_service: func create(), creating a new tinyurl for the user " + url.UserID + " and url - " + url.OriginalURL)
	key, err := getID()
	if err != nil {
		return nil, err
	}
	fmt.Println(strings.Trim(*key, "\""))
	url.TinyURL = "https://tinyurl.com/" + strings.Trim(*key, "\"")
	url.CreationDate = dateutils.GetNow()
	url.ExpirationDate = dateutils.GetExpiry()
	err = url.Save()

	if err != nil {
		if strings.Contains(err.Message, "already exists") {
			services.KeyService.SetKey(*key)
		}
		return nil, err
	}
	return &url, nil
}

func (u *urlService) GetURL(id string) (tinyurl.Urls, *errors.RestErr) {
	fmt.Println("Inside GetURL")
	url := tinyurl.Url{
		UserID: id,
	}
	result, getErr := url.List()
	if getErr != nil {
		return nil, getErr
	}
	return result, nil
}

func (u *urlService) DeleteURL(url tinyurl.Url) *errors.RestErr {
	if validateErr := url.Validate(); validateErr != nil {
		return validateErr
	}
	delErr := url.Delete()
	if delErr != nil {
		return delErr
	}
	return nil
}

func getID() (*string, *errors.RestErr) {
	zlogger.Info("url_Services : func getID(), trying to fetch the key present in the cache of app services")
	for {
		if res := services.KeyService.Get(); res != "" {
			zlogger.Info("url_Services : func getID(), unique key successfully found in local cache")
			return &res, nil
		}
		zlogger.Info("url_Services : func getID(), key is not present in cache, an internal API call is made to fetch key from kgs services")
		response := kgsRestClient.Get("/getkey")
		if response == nil || response.Response == nil {
			return nil, errors.NewInternalServerError("invalid rest client response when trying to fetch keys from kgs store")
		}
		if response.StatusCode > 299 {
			var restErr errors.RestErr
			err := json.Unmarshal(response.Bytes(), &restErr)
			if err != nil {
				return nil, errors.NewInternalServerError("Invalid error interface while trying to get key")
			}
			return nil, &restErr
		}
		zlogger.Info("url_Services : func getID(), returning the key received from the kgs services")
		keys := trimResult(response.Bytes())
		services.KeyService.Set(keys)
	}

}

func trimResult(inp []byte) []string {
	res := string(inp)
	res = strings.TrimSuffix(res, "]")
	res = strings.TrimPrefix(res, "[")
	return strings.Split(res, ",")
}
