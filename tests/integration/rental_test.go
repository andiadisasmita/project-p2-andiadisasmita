package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateRental(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	payload := `{"user_id":1,"boardgame_id":1,"rental_period":7}`
	req := httptest.NewRequest(http.MethodPost, "/rentals", bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test CreateRental Controller
	err = controllers.CreateRental(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Parse and Validate the Response
		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Rental created successfully", response["message"])
	}
}

func TestGetUserRentals(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	req := httptest.NewRequest(http.MethodGet, "/rentals", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test GetUserRentals Controller
	err = controllers.GetUserRentals(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse and Validate the Response
		var response []map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	}
}
