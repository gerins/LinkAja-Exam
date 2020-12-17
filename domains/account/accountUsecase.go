package account

import (
	"errors"
	"log"
)

// Usecase layer contains application business rules
type AccountUsecase struct {
	AccountRepo AccountRepositoryInterface
}

func NewAccountUsecase(repo AccountRepositoryInterface) *AccountUsecase {
	return &AccountUsecase{AccountRepo: repo}
}

type AccountUsecaseInterface interface {
	GetAccountInfo(accountNumber string) (*Account, error)
	ProcessingTransaction(transDetail *Transaction) error
}

func (u *AccountUsecase) GetAccountInfo(accountNumber string) (*Account, error) {
	result, err := u.AccountRepo.GetAccountInfo(accountNumber)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

// Validate transaction request before processing to Repository layer
func (u *AccountUsecase) ProcessingTransaction(transDetail *Transaction) error {
	// Validate input amount, cannot zero or negative number
	if transDetail.Amount < 1 {
		log.Println(transDetail, "invalid amount")
		return errors.New("Invalid Amount")
	}

	// Checking Sender Account avaibility & current balance
	senderAcc, err := u.AccountRepo.GetAccountInfo(transDetail.SenderAccount)
	if err != nil {
		log.Println(transDetail, "sender account not found")
		return errors.New("Sender account not found")
	} else if senderAcc.Balance < transDetail.Amount {
		log.Println(transDetail, "not enough balance")
		return errors.New("Not Enough Balance")
	}

	// Checking Receiver Account avaibility
	_, err = u.AccountRepo.GetAccountInfo(transDetail.ReceiverAccount)
	if err != nil {
		log.Println(transDetail, "receiver account not found")
		return errors.New("Receiver account not found")
	}

	// If everything good then process to Repository layer
	err = u.AccountRepo.ProcessingTransaction(transDetail)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
