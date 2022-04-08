package tinyurl

import (
	"fmt"
	"strings"
	"test3/hariprathap-hp/system_design/TinyURL/dataResources/postgresDB/urls_db"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"
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
		zlogger.Error("url_dao: func save(), db statement preparation failed with error : ", errors.NewError(err.Error()))
		return errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	//user_id, _ := strconv.Atoi(url.UserID)
	if _, insertErr := stmt.Exec(url.TinyURL, url.OriginalURL, url.CreationDate, url.ExpirationDate, url.UserID); insertErr != nil {
		if strings.Contains(insertErr.Error(), indexUniqueUserID) {
			zlogger.Error("url_dao: func save(), db statement preparation failed with error : ", errors.NewError("user already exists in db"))
			return errors.NewInternalServerError(fmt.Sprintf("user %s already exists", url.UserID))
		}
		zlogger.Error("url_dao: func save(), creation of url in db failed with error : ", errors.NewError(insertErr.Error()))
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user : %s", insertErr.Error()))
	}
	zlogger.Info("url_dao: func save(), url is successfully created and saved in db")
	return nil
}

func (url *Url) List() (Urls, *errors.RestErr) {
	stmt, err := urls_db.Client.Prepare(searchQuery)
	if err != nil {
		zlogger.Error("url_dao: func list(), db statement preparation failed with error : ", errors.NewError(err.Error()))
		return nil, errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	//user_id, _ := strconv.Atoi(url.UserID)
	rows, searchErr := stmt.Query(url.UserID)
	if searchErr != nil {
		zlogger.Error("url_dao: func list(), fetching the list of urls from db failed with error : ", errors.NewError(searchErr.Error()))
		return nil, errors.NewInternalServerError("fetching users from database failed")
	}
	defer rows.Close()

	results := make([]Url, 0)
	for rows.Next() {
		var res Url
		scanErr := rows.Scan(&res.TinyURL, &res.OriginalURL,
			&res.CreationDate, &res.ExpirationDate)
		if scanErr != nil {
			zlogger.Error("url_dao: func list(), scanning of results into url objects failed with error : ", errors.NewError(scanErr.Error()))
			return nil, errors.NewInternalServerError("failed during scanning result rows")
		}
		results = append(results, res)
	}
	zlogger.Info("url_dao: func list(), list of urls successfully fetched from db")
	return results, nil
}

func (url *Url) Delete() *errors.RestErr {
	stmt, err := urls_db.Client.Prepare(deleteQuery)
	if err != nil {
		return errors.NewInternalServerError("databse error")
	}
	defer stmt.Close()
	//user_id, _ := strconv.Atoi(url.UserID)
	if _, deleteErr := stmt.Exec(url.UserID, url.OriginalURL); deleteErr != nil {
		if strings.Contains(deleteErr.Error(), indexUniqueUserID) {
			fmt.Println("violates unique constraint")
			return errors.NewInternalServerError(fmt.Sprintf("user %s already exists", url.UserID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user : %s", deleteErr.Error()))
	}

	return nil
}
