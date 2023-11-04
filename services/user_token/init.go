package token

import (
	"golang-server-init/app"
)

var token service

func Init(app *app.App) {
	token = newService("Token", *app.Config)
	app.RegisterService(token)
}
