package logout

import (
	"golang-server-init/app"
)

var logout service

func Init(app *app.App) {
	logout = newService("Logout", *app.Config)
	app.RegisterService(logout)
}
