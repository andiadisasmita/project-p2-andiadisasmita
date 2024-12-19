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

func TestRegisterUser(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	payload := `{"username":"testuser","email":"testuser@example.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test RegisterUser Controller
	err = controllers.RegisterUser(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Parse and Validate the Response
		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "User registered successfully", response["message"])
	}
}

func TestLoginUser(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Ensure User Exists in Database
	// Add user creation logic here if necessary for test setup

	// Mock HTTP Request
	payload := `{"email":"testuser@example.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test LoginUser Controller
	err = controllers.LoginUser(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse and Validate the Response
		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "token")
	}
}
