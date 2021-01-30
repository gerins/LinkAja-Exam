package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestNewAccountController(t *testing.T) {
	db, _ := newDbMock()
	defer db.Close()

	// repo := NewAccountController(db)
	// assert.NotNil(t, repo)
}

func TestAccountController_HandleGetAccountInfo(t *testing.T) {
	db, _ := newDbMock()
	defer db.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/account/:id")
	c.SetParamNames("id")
	c.SetParamValues("555001")

	// h := NewAccountController(db)
	// if assert.NoError(t, h.HandleGetAccountInfo(c)) {
	// 	assert.Equal(t, http.StatusNotFound, rec.Code)
	// }
}

func TestAccountController_HandleProcessingTransaction(t *testing.T) {
	db, _ := newDbMock()
	defer db.Close()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/account/:id/transfer")
	c.SetParamNames("id")
	c.SetParamValues("555001")

	// h := NewAccountController(db)
	// if assert.NoError(t, h.HandleProcessingTransaction(c)) {
	// 	assert.Equal(t, http.StatusBadRequest, rec.Code)
	// }
}
