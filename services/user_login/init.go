package login

import (
	"golang-server-init/app"
)

var login service

func init() {

}

func Init(app *app.App) {
	login = newService("Login", *app.Config)
	app.RegisterService(login)
}
