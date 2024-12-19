package controllers

import (
	"net/http"
	"time"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
)

// GetUserRentals godoc
// @Summary Retrieve active rentals for the logged-in user
// @Description Returns a list of boardgames currently rented by the logged-in user
// @Tags rentals
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.RentalHistory
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /rentals [get]

// CreateRental godoc
// @Summary Rent a boardgame
// @Description Creates a rental record for a boardgame
// @Tags rentals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rental body models.CreateRentalRequest true "Rental details"
// @Success 201 {object} models.RentalHistory
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /rentals [post]

// UpdateRental godoc
// @Summary Update rental status
// @Description Updates the status of a rental record (e.g., mark as returned)
// @Tags rentals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Rental ID"
// @Param rental body models.UpdateRentalRequest true "Updated rental status"
// @Success 200 {object} models.RentalHistory
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /rentals/{id} [put]

// GetRentalHistory godoc
// @Summary Retrieve rental history for the logged-in user
// @Description Returns the past rentals of the logged-in user
// @Tags rentals
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.RentalHistory
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /rentals/history [get]

// GetUserRentals retrieves all active rentals for the logged-in user
func GetUserRentals(c echo.Context) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", err.Error()))
	}

	var rentals []models.RentalHistory
	if err := config.DB.Where("user_id = ? AND status != 'returned'", userID).Find(&rentals).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error retrieving rentals", err.Error()))
	}

	return c.JSON(http.StatusOK, rentals)
}

// CreateRental allows the user to rent a boardgame
func CreateRental(c echo.Context) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", err.Error()))
	}

	var input struct {
		StockID uint `json:"stock_id" validate:"required"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid input", err.Error()))
	}

	// Check stock availability
	var stock models.Stock
	if err := config.DB.First(&stock, input.StockID).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Stock not found", err.Error()))
	}
	if stock.Status != "warehouse" {
		return c.JSON(http.StatusConflict, utils.NewErrorResponse("Stock is not available for rental", ""))
	}

	// Create rental history
	rental := models.RentalHistory{
		UserID:     userID,
		StockID:    stock.ID,
		RentalDate: time.Now(),
		Status:     "reserved",
		RentalCost: 100.00, // Example rental cost, can be dynamic
	}
	if err := config.DB.Create(&rental).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error creating rental", err.Error()))
	}

	// Update stock status
	stock.Status = "with_user"
	if err := config.DB.Save(&stock).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error updating stock", err.Error()))
	}

	return c.JSON(http.StatusCreated, rental)
}

// UpdateRental updates the rental status (e.g., return the boardgame)
func UpdateRental(c echo.Context) error {
	var rental models.RentalHistory
	if err := config.DB.First(&rental, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Rental not found", err.Error()))
	}

	var input struct {
		Status string `json:"status" validate:"required"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid input", err.Error()))
	}

	rental.Status = input.Status
	if input.Status == "returned" {
		// Update stock status
		var stock models.Stock
		if err := config.DB.First(&stock, rental.StockID).Error; err == nil {
			stock.Status = "warehouse"
			config.DB.Save(&stock)
		}
	}

	if err := config.DB.Save(&rental).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error updating rental", err.Error()))
	}

	return c.JSON(http.StatusOK, rental)
}

// GetRentalHistory retrieves the rental history of a user
func GetRentalHistory(c echo.Context) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", err.Error()))
	}

	var history []models.RentalHistory
	if err := config.DB.Where("user_id = ?", userID).Find(&history).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error retrieving rental history", err.Error()))
	}

	return c.JSON(http.StatusOK, history)
}
