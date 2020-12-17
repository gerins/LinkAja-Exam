package main

import (
	"Linkaja/config"
	"Linkaja/config/database"
	"Linkaja/config/server"

	"github.com/labstack/echo"
)

func main() {
	// log.SetOutput(logging.InitRotateLogger())
	config.SetupEnvironmentVariable()

	srv := server.InitServer(echo.New(), database.InitDatabase())
	srv.InitRoutes().ApplyMiddleware().StartServer()
}
