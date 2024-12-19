package controllers

import (
	"net/http"
	"time"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
)

// GetPayments retrieves all payments for the logged-in user
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

// CreatePayment creates a new payment for a rental
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
