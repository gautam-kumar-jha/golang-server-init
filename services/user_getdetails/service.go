package getdetails

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-server-init/app"
	db "golang-server-init/app/database"
	es "golang-server-init/app/service"
	"log"
	"net/http"
	"strings"
)

type service struct {
	name        string
	config      app.Config
	dbProcessor *sql.DB
}

func newService(name string, config app.Config) service {
	return service{
		name:   name,
		config: config,
	}
}

func (getUserDetails service) GetHandler() es.Handler {
	return es.Handler{
		Method:      http.MethodGet,
		Path:        "api/v1/users/details",
		HandlerFunc: GetUserDetails(),
	}
}

func (getUserDetails service) Execute(req *http.Request) (es.ResponseEnvelope, int) {
	response := es.ResponseEnvelope{}

	userID, statusCode, err := validateRequest(req, getUserDetails.dbProcessor, getUserDetails.config.DatabaseConfig)
	if err != nil {
		response.Message = err.Error()
		return response, statusCode
	}

	// store data in database
	isSucess, result, err := processData(userID, getUserDetails.dbProcessor, getUserDetails.config.DatabaseConfig)
	response.IsSucess = isSucess

	if err != nil {
		response.Message = err.Error()
		return response, http.StatusInternalServerError
	}
	response.ResponseBody = result
	return response, http.StatusOK
}

func processData(userID string, DBProcessor *sql.DB, config db.Config) (bool, interface{}, error) {

	response := Response{}

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("database error")
	}

	defer DBProcessor.Close()

	addressJSON := ""
	address := Address{}
	cQuery := `SELECT * FROM UserProfile where userID = ? `
	err = DBProcessor.QueryRow(cQuery, userID).Scan(&response.UserID, &response.Name, &response.Email, &addressJSON)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("unable to find home address %s", err.Error())
		return false, response, fmt.Errorf("unable to find home address")
	}
	_ = json.Unmarshal([]byte(addressJSON), &address)
	response.Address = address

	response.MobileNo = strings.Split(response.UserID, "||")[1]
	response.Message = "Record Found."

	return true, response, nil
}
