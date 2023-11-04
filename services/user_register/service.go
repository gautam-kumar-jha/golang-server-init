package register

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-server-init/app"
	db "golang-server-init/app/database"
	es "golang-server-init/app/service"
	"io/ioutil"
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

func (registerService service) GetHandler() es.Handler {
	return es.Handler{
		Method:      http.MethodPost,
		Path:        "api/v1/users/register",
		HandlerFunc: RegisterUser(),
	}
}

func (registerService service) Execute(req *http.Request) (es.ResponseEnvelope, int) {

	request := Request{}
	response := es.ResponseEnvelope{}

	reqBody, err := ioutil.ReadAll(req.Body)
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

	msg, iv := request.validateRequest()
	if !iv {
		response.Message = msg
		response.IsSucess = false
		response.ResponseBody = nil
		return response, http.StatusBadRequest
	}

	// store data in database
	isSucess, result, err := processData(&request, registerService.dbProcessor, regUser.config.DatabaseConfig)
	response.IsSucess = isSucess

	if err != nil {
		response.Message = err.Error()
		return response, http.StatusInternalServerError
	}
	response.ResponseBody = result
	return response, http.StatusOK
}

func processData(req *Request, DBProcessor *sql.DB, config db.Config) (bool, interface{}, error) {

	response := Response{}

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("database error.")
	}
	defer DBProcessor.Close()

	mobileNo := strings.Trim(req.RequestBody.MobileNo, " ")
	userID := fmt.Sprintf("FM||%s", mobileNo)
	pswd := strings.Trim(req.RequestBody.Password, " ")
	name := strings.Trim(req.RequestBody.Name, " ")
	email := strings.Trim(req.RequestBody.Email, " ")

	// add data in user table
	userSqlStmt, err := DBProcessor.Prepare("INSERT INTO User (userID, mobileNo, pswd, isActive) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("error in user data insert statement %s", err.Error())
		return false, response, fmt.Errorf("user data not inserted")
	}
	defer userSqlStmt.Close()
	_, err = userSqlStmt.Exec(userID, mobileNo, pswd, "Y")
	if err != nil {
		log.Printf("error in user data insert %s", err.Error())
		return false, response, fmt.Errorf("user data not inserted")
	}

	// add data in profile table
	proSqlStmt, err := DBProcessor.Prepare("INSERT INTO UserProfile (userID, uName, uEmail) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("error in user data insert statement %s", err.Error())
		return false, response, fmt.Errorf("user data not inserted")
	}
	defer proSqlStmt.Close()
	_, err = proSqlStmt.Exec(userID, name, email)
	if err != nil {
		log.Printf("error in user data insert %s", err.Error())
		return false, response, fmt.Errorf("user data not inserted")
	}

	response.UserID = userID
	response.Message = "Sucessfully Registered."

	return true, response, nil
}
