package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockController struct{}

func (mc *MockController) RegisterUser(c echo.Context) error {
	// Simulate successful user registration
	email := c.FormValue("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid email",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user registered successfully",
	})
}

func TestMockController_RegisterUser(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	req.Form.Set("email", "testuser@example.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller := &MockController{}

	err := controller.RegisterUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "user registered successfully")
}
