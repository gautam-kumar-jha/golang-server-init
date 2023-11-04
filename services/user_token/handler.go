package token

import (
	"encoding/json"
	"net/http"
)

var Token = func() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, req *http.Request) {
		response, statusCode := token.Execute(req)
		generateResponse(responseWriter, statusCode, response)
	}
}

func generateResponse(rw http.ResponseWriter, StatusCode int, res interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(StatusCode)
	_ = json.NewEncoder(rw).Encode(res)
}
