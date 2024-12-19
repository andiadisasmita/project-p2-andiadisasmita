package controllers

import (
	"net/http"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
)

// GetReviews godoc
// @Summary Retrieve reviews for a boardgame
// @Description Returns a list of reviews for a specific boardgame
// @Tags reviews
// @Produce json
// @Param boardgame_id path int true "Boardgame ID"
// @Success 200 {array} models.Review
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /reviews/{boardgame_id} [get]

// CreateReview godoc
// @Summary Submit a review for a boardgame
// @Description Allows a user to submit a review for a boardgame they have rented
// @Tags reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param review body models.CreateReviewRequest true "Review details"
// @Success 201 {object} models.Review
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /reviews [post]

// GetReviews retrieves all reviews for a boardgame
func GetReviews(c echo.Context) error {
	var reviews []models.Review
	if err := config.DB.Where("boardgame_id = ?", c.Param("boardgame_id")).Find(&reviews).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error retrieving reviews", err.Error()))
	}

	return c.JSON(http.StatusOK, reviews)
}

// CreateReview allows a user to create a review for a boardgame
func CreateReview(c echo.Context) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Unauthorized", err.Error()))
	}

	var input models.Review
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid input", err.Error()))
	}
	input.UserID = userID

	if err := config.DB.Create(&input).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error creating review", err.Error()))
	}

	return c.JSON(http.StatusCreated, input)
}
