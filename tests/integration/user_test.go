package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Mock SendEmail function
var SendEmail func(to, subject, body string) error

func TestRegisterUser(t *testing.T) {
	e := echo.New()

	reqBody := `{"username":"testuser","email":"testuser@example.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader([]byte(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock SendEmail implementation
	emailSent := false
	SendEmail = func(to, subject, body string) error {
		if to == "" || subject == "" || body == "" {
			return errors.New("invalid email details")
		}
		emailSent = true
		return nil
	}

	err := controllers.RegisterUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully", response["message"])

	assert.True(t, emailSent, "Welcome email should be sent upon registration")
}

func TestLoginUser(t *testing.T) {
	e := echo.New()

	reqBody := `{"email":"testuser@example.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader([]byte(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controllers.LoginUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
}
