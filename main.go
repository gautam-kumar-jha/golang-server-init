package main

import (
	"golang-server-init/app"
	s "golang-server-init/services"
)

func main() {
	vapp := app.NewApp()
	s.Init(vapp)
	defer vapp.Start()
}

func SayHello() string {
	return "hello"
}
