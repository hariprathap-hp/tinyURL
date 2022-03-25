package urls

import (
	"fmt"
	"strconv"
	"strings"
	"test3/hariprathap-hp/system_design/TinyURL/dataResources/postgresDB/urls_db"
	"test3/hariprathap-hp/system_design/TinyURL/utils/errors"
)

const (
	indexUniqueUserID = "duplicate key value"
	insertQuery       = "insert into url (hash,originalurl,creationdate,expirationdate,userid) values ($1,$2,$3,$4,$5)"
)

func (url *Url) Save() *errors.RestErr {
	stmt, err := urls_db.Client.Prepare(insertQuery)
	if err != nil {
		return errors.NewInternalServerError("db query statement creation failed")
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

func (url *Url) List() *errors.RestErr {
	return nil
}
