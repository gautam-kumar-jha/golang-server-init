package service

import (
	"golang-server-init/app"
	getuser "golang-server-init/services/user_getdetails"
	login "golang-server-init/services/user_login"
	logout "golang-server-init/services/user_logout"
	register "golang-server-init/services/user_register"
	token "golang-server-init/services/user_token"
	updateprofile "golang-server-init/services/user_updateprofile"
)

// Init initiate all the services
func Init(app *app.App) {
	register.Init(app)
	getuser.Init(app)
	updateprofile.Init(app)
	login.Init(app)
	logout.Init(app)
	token.Init(app)
}
