package updateprofile

import (
	"context"
	"encoding/json"
	"net/http"
)

var SetAddress = func() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, req *http.Request) {
		ctx := context.TODO()
		response, statusCode := updateProfile.Execute(ctx, req)
		generateResponse(responseWriter, statusCode, response)
	}
}

func generateResponse(rw http.ResponseWriter, StatusCode int, res interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(StatusCode)
	_ = json.NewEncoder(rw).Encode(res)
}
