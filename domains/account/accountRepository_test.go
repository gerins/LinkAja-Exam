package account

import (
	"Linkaja/utils/query"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newDbMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAccountRepository_GetAccountInfo(t *testing.T) {
	db, mock := newDbMock()
	defer db.Close()

	acc := &Account{
		AccountNumber: "555003",
		CustomerName:  "GoLang",
		Balance:       20000,
	}

	repo := NewAccountRepository(db)
	rows := sqlmock.NewRows([]string{"ACCOUNT_NUMBER", "NAME", "BALANCE"}).
		AddRow(acc.AccountNumber, acc.CustomerName, acc.Balance)
	mock.ExpectQuery(query.ACCOUNT_INFO).WithArgs(acc.AccountNumber).WillReturnRows(rows)

	user, err := repo.GetAccountInfo(acc.AccountNumber)
	assert.Equal(t, "555003", user.AccountNumber)
	assert.Equal(t, "GoLang", user.CustomerName)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAccountRepository_ProcessingTransaction(t *testing.T) {
	db, mock := newDbMock()
	defer db.Close()

	trans := &Transaction{
		SenderAccount:   "555001",
		ReceiverAccount: "55502",
		Amount:          1000,
	}

	repo := NewAccountRepository(db)
	mock.ExpectBegin()
	mock.ExpectPrepare("UPDATE account SET BALANCE = BALANCE -").ExpectExec().WithArgs(trans.Amount, trans.SenderAccount).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("UPDATE account SET BALANCE = BALANCE +").ExpectExec().WithArgs(trans.Amount, trans.ReceiverAccount).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.ProcessingTransaction(trans)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestNewAccountRepository(t *testing.T) {
	db, _ := newDbMock()
	defer db.Close()

	repo := NewAccountRepository(db)
	assert.NotNil(t, repo)
}
