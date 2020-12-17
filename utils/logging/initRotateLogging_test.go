package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRotateLogger(t *testing.T) {
	rl := InitRotateLogger()

	assert.NotNil(t, rl)
	assert.NotNil(t, rl.CurrentFileName())
}
