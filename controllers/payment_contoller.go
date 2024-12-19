package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
)

// GetPayments godoc
// @Summary Retrieve payments for the logged-in user
// @Description Returns a list of payments made by the logged-in user
// @Tags payments
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Payment
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /payments [get]
func GetPayments(c echo.Context) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", err.Error()))
	}

	var payments []models.Payment
	if err := config.DB.Joins("JOIN rental_histories ON payments.rental_id = rental_histories.id").
		Where("rental_histories.user_id = ?", userID).
		Find(&payments).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error retrieving payments", err.Error()))
	}

	return c.JSON(http.StatusOK, payments)
}

// CreatePayment godoc
// @Summary Create a payment
// @Description Records a payment for a rental
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payment body models.CreatePaymentRequest true "Payment details"
// @Success 201 {object} models.Payment
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /payments [post]
func CreatePayment(c echo.Context) error {
	var input struct {
		RentalID uint    `json:"rental_id" validate:"required"`
		Amount   float64 `json:"amount" validate:"required"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid input", err.Error()))
	}

	payment := models.Payment{
		RentalID: input.RentalID,
		Amount:   input.Amount,
		Status:   "paid",
		PaidAt:   time.Now(),
	}
	if err := config.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error creating payment", err.Error()))
	}

	return c.JSON(http.StatusCreated, payment)
}

// CreateInvoice godoc
// @Summary Create a payment invoice
// @Description Generates a payment invoice for a rental using the Xendit API.
// @Tags payments
// @Accept json
// @Produce json
// @Param invoice body struct{RentalID uint; Amount float64; Description string} true "Invoice details"
// @Success 201 {object} map[string]string "Invoice URL and success message"
// @Failure 400 {object} utils.ErrorResponse "Invalid input"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /payments/invoice [post]
func CreateInvoice(c echo.Context) error {
	type InvoiceRequest struct {
		RentalID    uint    `json:"rental_id"`
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	req := new(InvoiceRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload", err.Error()))
	}

	invoiceURL, err := utils.CreateInvoice(req.RentalID, req.Amount, req.Description, os.Getenv("SUCCESS_REDIRECT_URL"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to create invoice", err.Error()))
	}

	email := "<USER_EMAIL>" // Retrieve from the user's context or rental details
	subject := "Payment Confirmation"
	body := "<h1>Payment Successful</h1><p>Thank you for your payment!</p>"

	if err := utils.SendEmail(email, subject, body); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to send payment confirmation email",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message":     "Invoice created successfully and payment confirmation email sent",
		"invoice_url": invoiceURL,
	})
}

// CheckInvoice godoc
// @Summary Check payment invoice status
// @Description Retrieves the status of a payment invoice from the Xendit API.
// @Tags payments
// @Produce json
// @Param invoice_id path string true "Invoice ID"
// @Success 200 {object} map[string]string "Invoice status"
// @Failure 404 {object} utils.ErrorResponse "Invoice not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /payments/status/{invoice_id} [get]
func CheckInvoice(c echo.Context) error {
	invoiceID := c.Param("invoice_id")
	if invoiceID == "" {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid Invoice ID", "Invoice ID is required"))
	}

	status, err := utils.CheckInvoiceStatus(invoiceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to retrieve invoice status", err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"invoice_id": invoiceID,
		"status":     status,
	})
}
