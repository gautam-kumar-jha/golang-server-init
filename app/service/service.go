package service

import (
	"context"
	"net/http"
)

type Service interface {
	GetHandler() Handler
	Execute(context context.Context, req *http.Request) (ResponseEnvelope, int)
}
