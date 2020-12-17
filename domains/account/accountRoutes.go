package account

import (
	"database/sql"

	"github.com/labstack/echo"
)

// Initializing API group routes
func InitAccountRoutes(mainRoute string, e *echo.Echo, db *sql.DB) {
	// Depedency Injection
	accountRepository := NewAccountRepository(db)
	accountUsecase := NewAccountUsecase(accountRepository)
	accountController := NewAccountController(accountUsecase)

	r := e.Group(mainRoute)
	r.GET("/:id", accountController.HandleGetAccountInfo)
	r.POST("/:senderId/transfer", accountController.HandleProcessingTransaction)
}
