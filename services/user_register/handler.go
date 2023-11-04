package register

import (
	"encoding/json"
	"net/http"
)

var RegisterUser = func() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, req *http.Request) {
		response, statusCode := regUser.Execute(req)
		generateResponse(responseWriter, statusCode, response)
	}
}

func generateResponse(rw http.ResponseWriter, StatusCode int, res interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(StatusCode)
	_ = json.NewEncoder(rw).Encode(res)
}
