package services

import (
	"test3/hariprathap-hp/system_design/TinyURL/domain/urls"
	"test3/hariprathap-hp/system_design/TinyURL/utils/dateutils"
	"test3/hariprathap-hp/system_design/TinyURL/utils/errors"
)

func CreateURL(url urls.Url) (*urls.Url, *errors.RestErr) {
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
