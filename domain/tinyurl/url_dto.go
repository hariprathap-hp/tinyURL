package tinyurl

import (
	"strconv"
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

type Urls []Url

func (url *Url) Validate() *errors.RestErr {
	url.OriginalURL = strings.TrimSpace(strings.ToLower(url.OriginalURL))
	if url.OriginalURL == "" {
		return errors.NewBadRequestError("Invalid URL Input")
	}
	url.UserID = strings.TrimSpace(strings.ToLower(url.UserID))
	if url.UserID == "" {
		return errors.NewBadRequestError("Invalid User ID, UserID can't be empty")
	}
	if _, err := strconv.Atoi(url.UserID); err != nil {
		return errors.NewBadRequestError("Invalid User ID, UserID is an integer")
	}
	return nil
}
