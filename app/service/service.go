package service

import "net/http"

type Service interface {
	GetHandler() Handler
	Execute(req *http.Request) (ResponseEnvelope, int)
}
