package tests

import (
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	response := utils.NewErrorResponse("Error occurred", "Details here")
	assert.Equal(t, "Error occurred", response.Message)
	assert.Equal(t, "Details here", response.Details)
}
