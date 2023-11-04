package database

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	SqlUserName      string `json:"sqlUserName"`
	SqlPassword      string `json:"sqlPassword"`
	DBType           string `json:"dbType"`
	ConnectionString string `json:"connectionString"`
	DatabaseName     string `json:"databaseName"`
	Host             string `json:"host"`
}

func (config *Config) SetConfig() {
	log.Println("Database Configuration :: Start")
	config.SqlUserName = os.Getenv("MYSQL_USERNAME")
	config.SqlPassword = os.Getenv("MYSQL_PASSWORD")
	config.DBType = os.Getenv("MYSQL_DBTYPE")
	config.DatabaseName = os.Getenv("MYSQL_DBNAME")
	config.Host = os.Getenv("MYSQL_HOST")
	config.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s", config.SqlUserName, config.SqlPassword, config.Host, config.DatabaseName)
	log.Println(config.ConnectionString)
	log.Println("Database Configuration :: Done")
}
