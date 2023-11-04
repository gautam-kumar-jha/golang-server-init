package updateprofile

import (
	"golang-server-init/app"
)

var updateProfile service

func Init(app *app.App) {
	updateProfile = newService("Update Profile", *app.Config)
	app.RegisterService(updateProfile)
}
