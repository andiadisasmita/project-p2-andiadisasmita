package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/routes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoutesSetup(t *testing.T) {
	// Set JWT_SECRET environment variable for testing
	os.Setenv("JWT_SECRET", "testsecret")
	defer os.Unsetenv("JWT_SECRET")

	// Initialize Echo instance
	e := echo.New()
	routes.SetupRoutes(e)

	// Test an unprotected route
	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	// Test a protected route without JWT token
	req = httptest.NewRequest(http.MethodGet, "/boardgames", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Expect unauthorized response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
