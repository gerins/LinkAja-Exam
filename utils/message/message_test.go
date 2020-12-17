package message

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespone(t *testing.T) {
	result := Respone("Success", http.StatusOK, true)
	assert.Equal(t, result.Code, 200)
	assert.Equal(t, result.Results, true)
}
