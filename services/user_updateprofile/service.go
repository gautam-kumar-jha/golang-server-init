package updateprofile

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
		Path:        "api/v1/users/updatedetails",
		HandlerFunc: SetAddress(),
	}
}

func (setAddress service) Execute(req *http.Request) (es.ResponseEnvelope, int) {

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

	msg, stCode, iv := request.isRequestDataValid(setAddress.dbProcessor, setAddress.config.DatabaseConfig)
	if !iv {
		response.Message = msg
		response.IsSucess = false
		return response, stCode
	}

	// store data in database
	isSucess, result, err := processData(&request, setAddress.dbProcessor, setAddress.config.DatabaseConfig)
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

	// set data model for request
	userID := strings.Trim(req.RequestBody.UserID, " ")
	userName := strings.Trim(req.RequestBody.Name, " ")
	email := strings.Trim(req.RequestBody.Email, " ")
	addData, err := json.Marshal(req.RequestBody.Address)
	if err != nil {
		log.Printf("unable to marshal address %s", err.Error())
		return false, response, fmt.Errorf("unable to marshal address.")
	}

	uPStmt, err := DBProcessor.Prepare("UPDATE UserProfile SET uName=?, uEmail=?, uAddress=? WHERE userID=?")
	if err != nil {
		panic(err.Error())
	}
	defer uPStmt.Close()
	sqlRes, err := uPStmt.Exec(userName, email, addData, userID)
	if err != nil {
		log.Printf("unable to update profile data %s", err.Error())
		return false, response, fmt.Errorf("unable to update profile data.")
	}

	_, err = sqlRes.RowsAffected()
	if err != nil {
		log.Printf("no user found %s", err.Error())
		return false, response, fmt.Errorf("no user found.")
	}
	defer uPStmt.Close()

	response.UserID = userID
	response.Message = "Profile Updated."

	return true, response, nil
}
