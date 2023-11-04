package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetDBProcessor(config Config) (*sql.DB, error) {
	db, err := sql.Open(config.DBType, config.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateDatabase(config Config) {

	db, err := GetDBProcessor(config)
	if err != nil {
		log.Printf("database error %s", err.Error())
	}
	defer db.Close()
  
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://./database/migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("database migration completed successfully")
}
