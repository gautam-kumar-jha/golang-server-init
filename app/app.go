package app

import (
	"golang-server-init/app/database"
	"golang-server-init/app/middleware"
	"golang-server-init/app/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type App struct {
	Config   *Config
	Router   *mux.Router
	Services []service.Service
}

func NewApp() *App {
	app := &App{
		Config:   &Config{},
		Router:   mux.NewRouter().StrictSlash(true),
		Services: []service.Service{},
	}
	app.Config.LoadConfig()
	app.registerMiddleware()
	database.MigrateDatabase(app.Config.DatabaseConfig)
	return app
}

func (application *App) RegisterService(service service.Service) {
	application.Services = append(application.Services, service)
}

func (application *App) startServices() {
	for _, service := range application.Services {
		handler := service.GetHandler()
		application.Router.HandleFunc(handler.Path, handler.HandlerFunc).Methods(handler.Method)
	}
}

func (app *App) registerMiddleware() {
	app.Router.Use(middleware.AuthorizeRequest(app.Router))
	app.Router.Use(middleware.LogRequest(app.Router))
}

func (app *App) Start() {
	app.startServices()
	log.Printf("Started %s :-   %s\n", app.Config.GenericConfig.AppName, app.Config.GenericConfig.Port)
	if err := http.ListenAndServe(":"+app.Config.GenericConfig.Port, app.Router); err != nil {
		log.Fatal(err)
	}
}
