package tinyurl

import (
	"fmt"
	"strconv"
	"test3/hariprathap-hp/system_design/TinyURL/dataResources/postgresDB/urls_db"
	"test3/hariprathap-hp/system_design/tinyURL/dataResources/cassandra"
	"test3/hariprathap-hp/system_design/tinyURL/logger"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
)

const (
	indexUniqueUserID  = "duplicate key value"
	queryinsertTinyURL = "INSERT into urls (tiny_url, original_url, creation_date, expiration_date, user_id) values (?,?,?,?,?)"
	querylistAllURLs   = "select "
	searchQuery        = " hash,originalurl,creationdate,expirationdate from url where user_id=$1"
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
	stmt, err := urls_db.Client.Prepare(searchQuery)
	if err != nil {
		logger.Error("error while trying to create db statement", err)
		return nil, errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	user_id, _ := strconv.Atoi(url.UserID)

	rows, searchErr := stmt.Query(user_id)
	if searchErr != nil {
		fmt.Println(searchErr)
		return nil, errors.NewInternalServerError("fetching users from database failed")
	}
	defer rows.Close()

	results := make([]Url, 0)
	for rows.Next() {
		var res Url
		scanErr := rows.Scan(&res.TinyURL, &res.OriginalURL,
			&res.CreationDate, &res.ExpirationDate)
		if scanErr != nil {
			return nil, errors.NewInternalServerError("failed during scanning result rows")
		}
		results = append(results, res)
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
