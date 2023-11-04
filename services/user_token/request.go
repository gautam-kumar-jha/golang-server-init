package token

import (
	"database/sql"
	db "golang-server-init/app/database"
	"log"
	"strings"
)

type RequestBody struct {
	MobileNo string `json:"mobileNo"`
	Password string `json:"password"`
}

type Request struct {
	RequestBody RequestBody `json:"requestBody,omitempty"`
}

func (req Request) validateRequest(DBProcessor *sql.DB, config db.Config) (string, string, bool) {
	if req.RequestBody.MobileNo == "" || req.RequestBody.Password == "" {
		return "Invalid Request.", "", false
	}
	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return "Invalid User.", "", false
	}
	defer DBProcessor.Close()

	var userTID, password string
	enQuery := `SELECT userID, pswd FROM User where mobileNo = ? `
	err = DBProcessor.QueryRow(enQuery, strings.Trim(req.RequestBody.MobileNo, " ")).Scan(&userTID, &password)
	if err != nil {
		log.Printf("invalid user %s", err.Error())
		return "Invalid User.", "", false
	}
	return strings.Trim(userTID, " "), strings.Trim(password, " "), true
}
