package service

import "net/http"

type Handler struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}
