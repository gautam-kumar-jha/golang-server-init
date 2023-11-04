package updateprofile

import (
	"database/sql"
	db "golang-server-init/app/database"
	"log"
	"net/http"
	"strings"
)

type RequestBody struct {
	UserID  string `json:"userID"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address struct {
		HouseNo     string `json:"houseNo,omitempty"`
		LandMark    string `json:"landMark,omitempty"`
		Area        string `json:"area,omitempty"`
		Thana       string `json:"thana,omitempty"`
		PostOffice  string `json:"postOffice,omitempty"`
		Town        string `json:"town,omitempty"`
		State       string `json:"state,omitempty"`
		AddressType string `json:"addressType"`
	} `json:"address"`
}

type Request struct {
	RequestBody RequestBody `json:"requestBody,omitempty"`
}

func (req Request) validateRequest() (string, bool) {

	userID := strings.Trim(req.RequestBody.UserID, " ")
	if userID == "" {
		return "user ID is missing from request", false
	}

	return "", true
}

func (req Request) isRequestDataValid(DBProcessor *sql.DB, config db.Config) (string, int, bool) {

	userID := strings.Trim(req.RequestBody.UserID, " ")
	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return "database error.", http.StatusInternalServerError, false
	}
	defer DBProcessor.Close()

	var ID string
	enQuery := `SELECT userID FROM UserProfile where userID = ?`
	err = DBProcessor.QueryRow(enQuery, userID).Scan(&ID)
	if err != nil && err == sql.ErrNoRows {
		log.Printf("user not found %s", err.Error())
		return "user not found.", http.StatusNotFound, false
	}

	return "", http.StatusOK, true
}
