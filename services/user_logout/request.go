package logout

import (
	"database/sql"
	db "golang-server-init/app/database"
	"log"
	"net/http"
	"strings"
)

func validateRequest(req *http.Request, DBProcessor *sql.DB, config db.Config) (string, string, bool) {

	authToken := strings.Trim(req.Header.Get("X-Req-Farmingam"), " ")
	userID := strings.Trim(req.URL.Query().Get("userID"), " ")

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return "database error.", "", false
	}
	defer DBProcessor.Close()

	var token string
	enQuery := `SELECT token FROM Token where userID = ? AND token = ?`
	err = DBProcessor.QueryRow(enQuery, userID, authToken).Scan(&token)
	if err != nil {
		log.Printf("invalid token %s", err.Error())
		return "Invalid Session", "", false
	}

	return token, userID, true
}
