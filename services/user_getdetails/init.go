package getdetails

import (
	"golang-server-init/app"
)

var getUserDetails service

// Init ...
func Init(app *app.App) {
	getUserDetails = newService("Get User Details", *app.Config)
	app.RegisterService(getUserDetails)
}
