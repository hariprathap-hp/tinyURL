package tinyurl

import (
	"strings"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	"time"
)

type Url struct {
	TinyURL        string    `json:"tinyURL"`
	OriginalURL    string    `json:"url"`
	CreationDate   time.Time `json:"creation"`
	ExpirationDate time.Time `json:"expiration"`
	UserID         string    `json:"user_id"`
}

type UrlList struct {
	TinyURL     string `json:"tinyURL"`
	OriginalURL string `json:"url"`
}

type Urls []Url
type UrlsList []UrlList

func (url *Url) Validate() *errors.RestErr {
	url.OriginalURL = strings.TrimSpace(strings.ToLower(url.OriginalURL))
	if url.OriginalURL == "" {
		return errors.NewBadRequestError("Invalid URL Input")
	}
	url.UserID = strings.TrimSpace(strings.ToLower(url.UserID))
	if url.UserID == "" {
		return errors.NewBadRequestError("Invalid User ID, UserID can't be empty")
	}
	return nil
}

func (url *Url) ValidateURL() *errors.RestErr {
	url.TinyURL = strings.TrimSpace(strings.ToLower(url.TinyURL))
	if url.TinyURL == "" {
		return errors.NewBadRequestError("Invalid URL Input")
	}
	return nil
}
