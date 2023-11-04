package logout

import (
	"context"
	"database/sql"
	"fmt"
	"golang-server-init/app"
	db "golang-server-init/app/database"
	es "golang-server-init/app/service"
	"log"
	"net/http"
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

func (login service) GetHandler() es.Handler {
	return es.Handler{
		Method:      http.MethodGet,
		Path:        "/api/v1/users/logout",
		HandlerFunc: Logout(),
	}
}

func (login service) Execute(context context.Context, req *http.Request) (es.ResponseEnvelope, int) {

	response := es.ResponseEnvelope{}

	msg, uID, iv := validateRequest(req, login.dbProcessor, login.config.DatabaseConfig)
	if !iv {
		response.Message = msg
		response.IsSucess = false
		response.ResponseBody = nil
		return response, http.StatusUnauthorized
	}

	// store data in database
	isSucess, result, err := processData(msg, uID, login.dbProcessor, login.config.DatabaseConfig)
	response.IsSucess = isSucess
	if err != nil {
		response.Message = err.Error()
		return response, http.StatusInternalServerError
	}
	response.ResponseBody = result
	return response, http.StatusOK
}

func processData(aToken string, uID string, DBProcessor *sql.DB, config db.Config) (bool, interface{}, error) {

	response := Response{}

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("Signout Failed.")
	}
	defer DBProcessor.Close()

	err = deleteRow(uID, aToken, DBProcessor)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("Signout Failed.")
	}

	response.Message = "Signed Out."
	response.ShowPage = "Home"

	return true, response, nil
}

func deleteRow(userID string, aToken string, DBProcessor *sql.DB) error {
	delStmt, err := DBProcessor.Prepare("DELETE from token WHERE token = ? AND userID = ?")
	if err != nil {
		log.Printf("unable to delete token %s", err.Error())
		return fmt.Errorf("token error.")
	}
	defer delStmt.Close()

	result, err := delStmt.Exec(aToken, userID)
	if err != nil {
		log.Printf("unable to exec delete statement %s", err.Error())
		return fmt.Errorf("token error.")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Printf("error statement %s", err.Error())
		return fmt.Errorf("token error.")
	}
	return nil
}
