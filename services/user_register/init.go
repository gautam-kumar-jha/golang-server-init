package register

import (
	"golang-server-init/app"
)

var regUser service

func Init(app *app.App) {
	regUser = newService("RegisterUser", *app.Config)
	app.RegisterService(regUser)
}
