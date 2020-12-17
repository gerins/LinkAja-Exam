package database

import (
	"Linkaja/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	databaseSource := config.DB_USER + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME

	// Open connection to Database
	db, err := sql.Open(config.DB_DRIVER, databaseSource)
	if err != nil {
		log.Fatal("Unable Connect to Database", err.Error())
	}

	// Checking database availbility
	if err := db.Ping(); err != nil {
		log.Fatal("Unable Connect to Database", err.Error())
	}

	log.Println("DataBase Successfully Connected")
	return db
}
