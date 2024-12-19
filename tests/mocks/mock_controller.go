package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockController struct {
	Service *MockService
}

func (mc *MockController) RegisterUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := mc.Service.CreateUser(username, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user registered successfully",
	})
}

func (mc *MockController) CreateBoardgame(c echo.Context) error {
	title := c.FormValue("title")

	err := mc.Service.CreateBoardgame(title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "boardgame created successfully",
	})
}

func (mc *MockController) ProcessPayment(c echo.Context) error {
	amountStr := c.QueryParam("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid amount format",
		})
	}

	err = mc.Service.ProcessPayment(amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "payment processed successfully",
	})
}

// Tests

func TestMockController_RegisterUser(t *testing.T) {
	e := echo.New()
	mockService := &MockService{UserExists: false}
	controller := &MockController{Service: mockService}

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader("username=test&password=pass123"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.RegisterUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "user registered successfully")
}

func TestMockController_ProcessPayment(t *testing.T) {
	e := echo.New()
	mockService := &MockService{PaymentError: nil}
	controller := &MockController{Service: mockService}

	req := httptest.NewRequest(http.MethodGet, "/payment?amount=100", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.ProcessPayment(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "payment processed successfully")
}
