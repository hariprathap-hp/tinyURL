package services

import (
	"fmt"
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	"test3/hariprathap-hp/system_design/tinyURL/utils/dateutils"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"

	"github.com/bwmarrin/snowflake"
)

var (
	UrlServices urlServicesInterface = &urlService{}
)

type urlService struct{}

type urlServicesInterface interface {
	CreateURL(tinyurl.Url) (*tinyurl.Url, *errors.RestErr)
	GetURL(string) (tinyurl.Urls, *errors.RestErr)
	DeleteURL(tinyurl.Url) *errors.RestErr
}

func (u *urlService) CreateURL(url tinyurl.Url) (*tinyurl.Url, *errors.RestErr) {
	if validateErr := url.Validate(); validateErr != nil {
		return nil, validateErr
	}

	url.TinyURL = "https://tinyurl.com/" + getID()
	fmt.Println(url.TinyURL)
	url.CreationDate = dateutils.GetNow()
	url.ExpirationDate = dateutils.GetExpiry()
	err := url.Save()

	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (u *urlService) GetURL(id string) (tinyurl.Urls, *errors.RestErr) {
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

func getID() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Generate a snowflake ID.
	id := node.Generate()
	result := id.Base36()[3:]
	return result
}
