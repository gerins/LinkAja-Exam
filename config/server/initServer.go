package server

import (
	"Linkaja/config"
	"Linkaja/domains/account"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Echo *echo.Echo
	DB   *sql.DB
}

func InitServer(e *echo.Echo, db *sql.DB) *Server {
	return &Server{Echo: e, DB: db}
}

// Initializing API routes
func (s *Server) InitRoutes() *Server {
	account.InitAccountRoutes("/account", s.Echo, s.DB)

	s.Echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	return s
}

// Applying middleware to the web server
func (s *Server) ApplyMiddleware() *Server {
	s.Echo.Use(middleware.CORS())
	s.Echo.Use(middleware.Logger())
	s.Echo.Use(middleware.Recover())

	return s
}

// Starting web server
func (s *Server) StartServer() {
	serverDomain := config.MAIN_SERVER_HOST + ":" + config.MAIN_SERVER_PORT
	hosts := map[string]*Server{
		`api.` + serverDomain: {Echo: s.Echo},
	}

	s.Echo.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(c.Response(), req)
		}

		return
	})

	log.Println(`Server start at http://api.` + serverDomain)
	s.Echo.Logger.Fatal(s.Echo.Start(serverDomain))
}
