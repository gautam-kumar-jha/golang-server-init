package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-server-init/app"
	db "golang-server-init/app/database"
	es "golang-server-init/app/service"
	"golang-server-init/utils"
	"io/ioutil"
	"log"
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

func (login service) GetHandler() es.Handler {
	return es.Handler{
		Method:      http.MethodPost,
		Path:        "api/v1/users/login",
		HandlerFunc: Login(),
	}
}

func (login service) Execute(req *http.Request) (es.ResponseEnvelope, int) {
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

	msg, status, iv := request.validateRequest(login.dbProcessor, login.config.DatabaseConfig)
	if !iv {
		response.Message = msg
		response.IsSucess = false
		response.ResponseBody = nil
		return response, status
	}

	// store data in database
	isSucess, result, err := processData(msg, login.dbProcessor, login.config.DatabaseConfig)
	response.IsSucess = isSucess
	if err != nil {
		response.Message = err.Error()
		return response, http.StatusInternalServerError
	}
	response.ResponseBody = result
	return response, http.StatusOK
}

func processData(message string, DBProcessor *sql.DB, config db.Config) (bool, interface{}, error) {

	response := Response{}
	uID := strings.Split(message, "||FARM||")[1]
	token := strings.Split(message, "||FARM||")[0]

	DBProcessor, err := db.GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("Login Failed.")
	}
	defer DBProcessor.Close()

	aToken := utils.EncryptMessage(message + "&&&&&&&&&&|||||||||||&Farmingam||||||" + message + time.Now().String())

	err = updateRow(token, aToken, DBProcessor)
	if err != nil {
		log.Printf("database error %s", err.Error())
		return false, response, fmt.Errorf("Login Failed.")
	}

	displayPage := "Show Profile"
	isCartAvailable := isDataInCart(uID, DBProcessor)
	if isCartAvailable {
		displayPage = "Show ViewCart"
	}

	response.AuthToken = aToken
	response.UserID = uID
	response.Message = "Login Successed"
	response.ShowPage = displayPage

	return true, response, nil
}

func updateRow(token string, aToken string, DBProcessor *sql.DB) error {
	delStmt, err := DBProcessor.Prepare("UPDATE token SET token = ?, tokenTime = ? WHERE token=?")
	if err != nil {
		log.Printf("unable to delete token %s", err.Error())
		return fmt.Errorf("token error.")
	}
	defer delStmt.Close()

	result, err := delStmt.Exec(aToken, time.Now().String(), token)
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

func isDataInCart(userID string, DBProcessor *sql.DB) bool {
	mNo := strings.Split(userID, "||")[1]
	is := false
	var cCount int
	cQuery := `SELECT COUNT(mobileNo) FROM Cart where mobileNo = ? `
	err := DBProcessor.QueryRow(cQuery, mNo).Scan(&cCount)
	if err != nil {
		log.Printf("unable to fetch cart count %s", err.Error())
	}
	if cCount >= 1 {
		is = true
	}

	return is
}
