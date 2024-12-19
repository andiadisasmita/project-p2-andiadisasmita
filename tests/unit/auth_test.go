package tests

import (
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := utils.GenerateJWT(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
	// Generate a test JWT
	token, _ := utils.GenerateJWT(1)

	// Validate the token
	parsedToken, err := utils.ValidateJWT(token)
	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)
}

func TestHashPassword(t *testing.T) {
	password := "securepassword"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	// Verify the hash
	assert.True(t, utils.CheckPasswordHash(password, hash))
}
