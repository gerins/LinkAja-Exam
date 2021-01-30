package account

import (
	"Linkaja/utils/query"
	"database/sql"
	"errors"
	"fmt"
)

// Data access layer
type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

type AccountRepositoryInterface interface {
	GetAccountInfo(accountNumber string) (*Account, error)
	ProcessingTransaction(transDetail *Transaction) error
}

func (r *AccountRepository) GetAccountInfo(accountNumber string) (*Account, error) {
	row := r.db.QueryRow(query.ACCOUNT_INFO, accountNumber)

	fmt.Println(`Repository Running`)

	account := new(Account)
	if err := row.Scan(&account.AccountNumber, &account.CustomerName, &account.Balance); err != nil {
		return nil, errors.New("Account Not Found")
	}

	return account, nil
}

// Processing transaction to database
func (r *AccountRepository) ProcessingTransaction(transDetail *Transaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Preparing for updating Sender balance, if error happen, immediately return and trigger Rollback
	updateSenderBalance, err := tx.Prepare(query.UPDATE_SENDER_BALANCE)
	if err != nil {
		return err
	}
	defer updateSenderBalance.Close()
	// Execute the prepared statement
	if _, err = updateSenderBalance.Exec(transDetail.Amount, transDetail.SenderAccount); err != nil {
		return err
	}

	// Then prepare updating Receiver balance
	updateReceiverBalance, err := tx.Prepare(query.UPDATE_RECEIVER_BALANCE)
	if err != nil {
		return err
	}
	defer updateReceiverBalance.Close()
	// Execute the prepared statement
	if _, err = updateReceiverBalance.Exec(transDetail.Amount, transDetail.ReceiverAccount); err != nil {
		return err
	}

	// If everything good then we can safely return and commit the transaction
	return tx.Commit()
}
