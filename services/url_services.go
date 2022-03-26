package services

import (
	"test3/hariprathap-hp/system_design/tinyURL/domain/tinyurl"
	"test3/hariprathap-hp/system_design/tinyURL/utils/dateutils"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
)

func CreateURL(url tinyurl.Url) (*tinyurl.Url, *errors.RestErr) {
	if validateErr := url.Validate(); validateErr != nil {
		return nil, validateErr
	}
	url.TinyURL = "hashed#232112"
	url.CreationDate = dateutils.GetNow()
	url.ExpirationDate = dateutils.GetExpiry()
	err := url.Save()

	if err != nil {
		return nil, err
	}
	return &url, nil
}

func GetURL(id string) (tinyurl.Urls, *errors.RestErr) {
	url := tinyurl.Url{
		UserID: id,
	}
	result, getErr := url.List()
	if getErr != nil {
		return nil, getErr
	}
	return result, nil
}
