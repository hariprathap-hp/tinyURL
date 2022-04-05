package tinyurl

import (
	"fmt"
	"strconv"
	"strings"
	"test3/hariprathap-hp/system_design/TinyURL/dataResources/postgresDB/urls_db"
	"test3/hariprathap-hp/system_design/tinyURL/logger"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
)

const (
	indexUniqueUserID = "duplicate key value"
	insertQuery       = "insert into url (hash,originalurl,creationdate,expirationdate,userid) values ($1,$2,$3,$4,$5)"
	searchQuery       = "select hash,originalurl,creationdate,expirationdate from url where userid=$1"
	deleteQuery       = "delete from url where userid=$1 and originalurl=$2"
)

func (url *Url) Save() *errors.RestErr {
	stmt, err := urls_db.Client.Prepare(insertQuery)
	if err != nil {
		logger.Error("error while trying to create db statement", err)
		return errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	user_id, _ := strconv.Atoi(url.UserID)
	if _, insertErr := stmt.Exec(url.TinyURL, url.OriginalURL, url.CreationDate, url.ExpirationDate, user_id); insertErr != nil {
		if strings.Contains(insertErr.Error(), indexUniqueUserID) {
			fmt.Println("violates unique constraint")
			return errors.NewInternalServerError(fmt.Sprintf("user %s already exists", url.UserID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user : %s", insertErr.Error()))
	}
	return nil
}

func (url *Url) List() (Urls, *errors.RestErr) {
	fmt.Println("Inside List URLs")
	stmt, err := urls_db.Client.Prepare(searchQuery)
	if err != nil {
		logger.Error("error while trying to create db statement", err)
		return nil, errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	user_id, _ := strconv.Atoi(url.UserID)
	fmt.Println(user_id)
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
	fmt.Println(results)
	return results, nil
}

func (url *Url) Delete() *errors.RestErr {
	stmt, err := urls_db.Client.Prepare(deleteQuery)
	if err != nil {
		logger.Error("error while trying to create db statement", err)
		return errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	user_id, _ := strconv.Atoi(url.UserID)
	if _, deleteErr := stmt.Exec(user_id, url.OriginalURL); deleteErr != nil {
		if strings.Contains(deleteErr.Error(), indexUniqueUserID) {
			fmt.Println("violates unique constraint")
			return errors.NewInternalServerError(fmt.Sprintf("user %s already exists", url.UserID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user : %s", deleteErr.Error()))
	}

	return nil
}
