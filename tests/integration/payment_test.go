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

func TestCreatePayment(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	payload := `{"rental_id":1,"amount":100.00}`
	req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test CreatePayment Controller
	err = controllers.CreatePayment(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Parse and Validate the Response
		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Payment created successfully", response["message"])
	}
}

func TestGetPayments(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	req := httptest.NewRequest(http.MethodGet, "/payments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test GetPayments Controller
	err = controllers.GetPayments(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse and Validate the Response
		var response []map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	}
}
