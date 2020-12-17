package server

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var srv *Server

func TestInitServer(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	srv = InitServer(echo.New(), db)
	assert.NotNil(t, srv)
}

func TestServer_ApplyMiddleware(t *testing.T) {
	result := srv.ApplyMiddleware()

	assert.NotNil(t, result)
	assert.Equal(t, srv, result)
}

func TestServer_InitRoutes(t *testing.T) {
	result := srv.InitRoutes()

	assert.NotNil(t, result)
	assert.Equal(t, srv, result)
}
