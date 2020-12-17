package config

import (
	"log"
	"os"
)

var MAIN_SERVER_HOST,
	MAIN_SERVER_PORT,
	DB_DRIVER,
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_NAME string

func SetupEnvironmentVariable() {
	MAIN_SERVER_HOST = os.Getenv("MAIN_SERVER_HOST")
	MAIN_SERVER_PORT = os.Getenv("MAIN_SERVER_PORT")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")

	log.Println("---------------ENVIRONMENT VARIABLES---------------")
	log.Println(`MAIN_SERVER_HOST =`, MAIN_SERVER_HOST)
	log.Println(`MAIN_SERVER_PORT =`, MAIN_SERVER_PORT)
	log.Println(`DB_DRIVER =`, DB_DRIVER)
	log.Println(`DB_USER =`, DB_USER)
	log.Println(`DB_PASSWORD =`, DB_PASSWORD)
	log.Println(`DB_HOST =`, DB_HOST)
	log.Println(`DB_PORT =`, DB_PORT)
	log.Println(`DB_NAME =`, DB_NAME)
	log.Println("---------------------------------------------------")
}
