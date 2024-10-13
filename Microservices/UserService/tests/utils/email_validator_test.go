package utils_test

import (
	"UserService/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	validEmail := "test@example.com"
	invalidEmail := "test@example"

	assert.True(t, utils.IsValidEmail(validEmail))
	assert.False(t, utils.IsValidEmail(invalidEmail))
}
