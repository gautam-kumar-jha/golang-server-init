package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	db "golang-server-init/app/database"
	"log"
	"net/http"
)

func EncryptMessage(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	hashBytes := h.Sum(nil)
	eStr := hex.EncodeToString(hashBytes)
	return eStr
}

func ValidateToken(t string, userID string, DBProcessor *sql.DB, config db.Config) (bool, int, error) {

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, http.StatusInternalServerError, errors.New("database error")
	}
	defer DBProcessor.Close()

	var token string
	enQuery := `SELECT token FROM Token where userID = ? AND token = ?`
	err = DBProcessor.QueryRow(enQuery, userID, t).Scan(&token)
	if err != nil {
		log.Printf("token :  %s", err.Error())
		return false, http.StatusUnauthorized, errors.New("token not associated with this user")
	}

	return true, http.StatusOK, nil
}
