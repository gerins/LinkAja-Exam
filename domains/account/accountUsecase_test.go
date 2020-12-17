package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var accUsecase AccountUsecaseInterface

func TestNewAccountUsecase(t *testing.T) {
	// db, _ := newDbMock()
	// defer db.Close()

	// accUsecase = NewAccountUsecase(db)
	// assert.NotNil(t, accUsecase)
}

func TestAccountUsecase_GetAccountInfo(t *testing.T) {
	result, err := accUsecase.GetAccountInfo("123")
	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestAccountUsecase_ProcessingTransaction(t *testing.T) {
	for _, tt := range []struct {
		name    string
		value   int
		wantErr bool
	}{
		{
			name:    "Test Transfer Negatif Amount",
			value:   -10000,
			wantErr: true,
		},
		{
			name:    "Test Transfer 0 Amount",
			value:   0,
			wantErr: true,
		},
	} {
		err := accUsecase.ProcessingTransaction(&Transaction{Amount: tt.value})
		if !tt.wantErr {
			assert.NoError(t, err, tt.name)
		} else {
			assert.Error(t, err, tt.name)
		}
	}
}
