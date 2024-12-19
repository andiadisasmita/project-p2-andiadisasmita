package tests

import (
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/stretchr/testify/assert"
)

func TestInitializeDatabase(t *testing.T) {
	// Attempt to initialize the database
	err := config.InitializeDatabase()

	// Validate that no error occurred during initialization
	assert.NoError(t, err, "Database initialization should not return an error")

	// Validate that the global DB variable is set
	assert.NotNil(t, config.DB, "Global DB variable should not be nil")
}
