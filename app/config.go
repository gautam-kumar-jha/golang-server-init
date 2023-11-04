package app

import (
	appDB "golang-server-init/app/database"
	"os"
)

type GenericConfig struct {
	AppName string
	Port    string
	Env     string
}

func (gConfig *GenericConfig) LoadGenericConfig() error {
	gConfig.AppName = os.Getenv("APP_NAME")
	gConfig.Port = os.Getenv("APP_PORT")
	gConfig.Env = os.Getenv("ENV")
	return nil
}

type Config struct {
	DatabaseConfig appDB.Config
	GenericConfig  GenericConfig
}

func (config *Config) LoadConfig() {
	config.GenericConfig.LoadGenericConfig()
	config.DatabaseConfig.SetConfig()
}
