package tinyurl

import (
	"fmt"
	"strconv"
	"test3/hariprathap-hp/system_design/tinyURL/dataResources/cassandra"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	"time"
)

const (
	queryinsertTinyURL = "INSERT into urls (tiny_url, original_url, creation_date, expiration_date, user_id) values (?,?,?,?,?)"
	querylistAllURLs   = "select tiny_url, original_url, creation_date, expiration_date from urls where user_id=? ALLOW FILTERING"
	querydeleteTinyURL = "delete from urls where user_id=? and original_url=?"
)

func (url *Url) Save() *errors.RestErr {
	user_id, _ := strconv.Atoi(url.UserID)
	if err := cassandra.GetSession().Query(queryinsertTinyURL, url.TinyURL, url.OriginalURL, url.CreationDate, url.ExpirationDate, user_id).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (url *Url) List() (Urls, *errors.RestErr) {

	user_id, _ := strconv.Atoi(url.UserID)
	fmt.Println("user_id list is --", user_id)

	var results Urls
	m := map[string]interface{}{}

	iter := cassandra.GetSession().Query(querylistAllURLs, user_id).Iter()

	for iter.MapScan(m) {
		results = append(results, Url{
			TinyURL:        m["tiny_url"].(string),
			OriginalURL:    m["original_url"].(string),
			CreationDate:   m["creation_date"].(time.Time),
			ExpirationDate: m["expiration_date"].(time.Time),
		})
		m = map[string]interface{}{}
	}
	if err := iter.Close(); err != nil {
		fmt.Println(err.Error())
		return nil, errors.NewInternalServerError(err.Error())
	}
	return results, nil
}

func (url *Url) Delete() *errors.RestErr {
	user_id, _ := strconv.Atoi(url.UserID)
	if err := cassandra.GetSession().Query(querydeleteTinyURL, user_id, url.OriginalURL).Exec(); err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
