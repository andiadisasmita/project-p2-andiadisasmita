package controllers

import (
	"net/http"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
)

// GetBoardgames godoc
// @Summary Retrieve all boardgames
// @Description Returns a list of all available boardgames
// @Tags boardgames
// @Produce json
// @Success 200 {array} models.Boardgame
// @Failure 500 {object} utils.ErrorResponse
// @Router /boardgames [get]

// GetBoardgameByID godoc
// @Summary Retrieve a specific boardgame by ID
// @Description Returns details of a boardgame identified by its ID
// @Tags boardgames
// @Produce json
// @Param id path int true "Boardgame ID"
// @Success 200 {object} models.Boardgame
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /boardgames/{id} [get]

// GetBoardgames retrieves all boardgames
func GetBoardgames(c echo.Context) error {
	var boardgames []models.Boardgame
	if err := config.DB.Find(&boardgames).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error retrieving boardgames", err.Error()))
	}
	return c.JSON(http.StatusOK, boardgames)
}

// GetBoardgameByID retrieves a single boardgame by ID
func GetBoardgameByID(c echo.Context) error {
	var boardgame models.Boardgame
	if err := config.DB.First(&boardgame, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Boardgame not found", err.Error()))
	}
	return c.JSON(http.StatusOK, boardgame)
}
