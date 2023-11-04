package login

import (
	"database/sql"
	db "golang-server-init/app/database"
	"golang-server-init/utils"
	"net/http"
	"strings"
)

type RequestBody struct {
	UserID string `json:"userID"`
	Token  string `json:"token"`
}

type Request struct {
	RequestBody RequestBody `json:"requestBody,omitempty"`
}

func (req Request) validateRequest(DBProcessor *sql.DB, config db.Config) (string, int, bool) {

	userID := strings.Trim(req.RequestBody.UserID, " ")
	token := strings.Trim(req.RequestBody.Token, " ")

	if userID == "" {
		return "userID is missing from request.", http.StatusBadRequest, false
	}
	if token == "" {
		return "token is missing from request..", http.StatusBadRequest, false
	}

	isTokenValid, status, err := utils.ValidateToken(token, userID, DBProcessor, config)
	if !isTokenValid && err != nil {
		return err.Error(), status, false
	}

	return token + "||FARM||" + userID, http.StatusOK, true
}
