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

// Mock CreateInvoice function
var CreateInvoice func(rentalID uint, amount float64, description, callbackURL string) (string, error)

// Mock CheckInvoiceStatus function
var CheckInvoiceStatus func(invoiceID string) (string, error)

func TestCreateInvoice(t *testing.T) {
	e := echo.New()

	reqBody := `{"rental_id":1,"amount":100.00,"description":"Rental for Chess Game"}`
	req := httptest.NewRequest(http.MethodPost, "/payments/invoice", bytes.NewReader([]byte(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock Xendit invoice creation
	CreateInvoice = func(rentalID uint, amount float64, description, callbackURL string) (string, error) {
		if rentalID != 1 || amount != 100.00 || description == "" || callbackURL == "" {
			return "", errors.New("invalid inputs")
		}
		return "https://checkout.xendit.co/invoice/12345", nil
	}

	err := controllers.CreateInvoice(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invoice created successfully", response["message"])
	assert.Contains(t, response["invoice_url"], "https://checkout.xendit.co/invoice/")
}

func TestCheckInvoiceStatus(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/payments/status/12345", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("invoice_id")
	c.SetParamValues("12345")

	// Mock Xendit invoice status check
	CheckInvoiceStatus = func(invoiceID string) (string, error) {
		if invoiceID != "12345" {
			return "", errors.New("invalid invoice ID")
		}
		return "PAID", nil
	}

	err := controllers.CheckInvoice(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "PAID", response["status"])
}
