package token

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-server-init/app"
	db "golang-server-init/app/database"
	es "golang-server-init/app/service"
	"golang-server-init/utils"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
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

func (registerService service) GetHandler() es.Handler {
	return es.Handler{
		Method:      http.MethodPost,
		Path:        "api/v1/users/token",
		HandlerFunc: Token(),
	}
}

func (token service) Execute(req *http.Request) (es.ResponseEnvelope, int) {
	request := Request{}
	response := es.ResponseEnvelope{}
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		response.IsSucess = false
		response.Message = "can not read request."
		response.ResponseBody = nil
		return response, http.StatusInternalServerError
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		response.IsSucess = false
		response.Message = "can not unmarshal request."
		response.ResponseBody = nil
		return response, http.StatusInternalServerError
	}
	uID, pswd, iv := request.validateRequest(token.dbProcessor, token.config.DatabaseConfig)
	if !iv {
		response.Message = uID
		response.IsSucess = false
		response.ResponseBody = nil
		return response, http.StatusBadRequest
	}

	// store data in database
	isSucess, result, err := processData(uID, pswd, token.dbProcessor, token.config.DatabaseConfig)
	response.IsSucess = isSucess
	if err != nil {
		response.Message = err.Error()
		return response, http.StatusInternalServerError
	}
	response.ResponseBody = result
	return response, http.StatusOK
}

func processData(userID string, password string, DBProcessor *sql.DB, config db.Config) (bool, interface{}, error) {

	response := Response{}

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("Token Not Generated.")
	}
	defer DBProcessor.Close()

	err = deleteRow(userID, DBProcessor)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("Token Not Generated.")
	}

	token := getToken(userID, password)
	adSqlStmt, err := DBProcessor.Prepare("INSERT INTO Token (userID, token, tokenTime) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("error in user data insert statement %s", err.Error())
		return false, response, fmt.Errorf("user data not inserted")
	}
	defer adSqlStmt.Close()

	_, err = adSqlStmt.Exec(userID, token, time.Now().String())
	if err != nil {
		log.Printf("error in token data insert %s", err.Error())
		return false, response, fmt.Errorf("token data not inserted")
	}

	response.Token = token
	response.UserID = userID
	response.Message = "Success"

	return true, response, nil
}

func deleteRow(userID string, DBProcessor *sql.DB) error {

	delStmt, err := DBProcessor.Prepare("DELETE from token WHERE userID = ?")
	if err != nil {
		log.Printf("unable to delete token %s", err.Error())
		return fmt.Errorf("token error.")
	}
	defer delStmt.Close()

	_, err = delStmt.Exec(userID)
	if err != nil {
		log.Printf("unable to exec delete statement %s", err.Error())
		return fmt.Errorf("token error.")
	}

	return nil
}

func getToken(uid string, pswd string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000)
	tToken := fmt.Sprintf("%d|$|%s||F||A$$R||M$$IN$$GA||M %sD||A||T||E%s", randomNum, uid, pswd, time.Now().String())
	token := strings.Trim(utils.EncryptMessage(tToken), " ")
	return token
}
