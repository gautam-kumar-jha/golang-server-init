package getdetails

import (
	"database/sql"
	"errors"
	db "golang-server-init/app/database"
	"golang-server-init/utils"
	"net/http"
	"strings"
)

func validateRequest(req *http.Request, DBProcessor *sql.DB, config db.Config) (string, int, error) {
	status := http.StatusBadRequest
	authToken := strings.Trim(req.Header.Get("X-Req-Farmingam"), " ")
	if authToken == "" {
		return "", status, errors.New("header token is missing from request")
	}

	userID := strings.Trim(req.URL.Query().Get("userID"), " ")
	if userID == "" {
		return "", status, errors.New("user ID is missing from request")
	}

	isTokenValid, status, err := utils.ValidateToken(authToken, userID, DBProcessor, config)
	if !isTokenValid && err != nil {
		return userID, status, err
	}
	status = http.StatusOK
	return userID, status, nil
}
