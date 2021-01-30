package account

import (
	"Linkaja/utils/message"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Controller layer for receiving and serializing input from user
type AccountController struct {
	AccountUsecase AccountUsecaseInterface
}

func NewAccountController(useCase AccountUsecaseInterface) *AccountController {
	return &AccountController{AccountUsecase: useCase}
}

func (c *AccountController) HandleGetAccountInfo(e echo.Context) error {
	accountNumber := e.Param("id")

	fmt.Println(`Controller Running`)

	result, err := c.AccountUsecase.GetAccountInfo(accountNumber)
	if err != nil {
		return e.JSON(http.StatusNotFound, message.Respone("Get Account Info Failed", http.StatusNotFound, err.Error()))
	}

	return e.JSON(http.StatusOK, result)
}

// Serializing user input then send to usecase layer
func (c *AccountController) HandleProcessingTransaction(e echo.Context) error {
	transDetail := &Transaction{SenderAccount: e.Param("senderId")}

	if err := e.Bind(transDetail); err != nil {
		return e.JSON(http.StatusBadRequest, message.Respone("Transaction Failed", http.StatusBadRequest, err.Error()))
	}

	// After serializing input from user the proceed to Usecase layer
	err := c.AccountUsecase.ProcessingTransaction(transDetail)
	if err != nil {
		return e.JSON(http.StatusBadRequest, message.Respone("Transaction Failed", http.StatusBadRequest, err.Error()))
	}

	return e.JSON(http.StatusCreated, nil)
}
