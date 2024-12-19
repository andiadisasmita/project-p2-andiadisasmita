package controllers

import (
	"net/http"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/models"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Allows a new user to register by providing email, password, and deposit amount
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User registration details"
// @Success 201 {object} echo.Map
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /users/register [post]

// LoginUser godoc
// @Summary Login a user
// @Description Authenticates a user and provides a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login credentials"
// @Success 200 {object} echo.Map
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /users/login [post]

// RegisterUser handles user registration
func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid input", err.Error()))
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error hashing password", err.Error()))
	}
	user.Password = hashedPassword

	// Save user
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error creating user", err.Error()))
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully"})
}

// LoginUser handles user login
func LoginUser(c echo.Context) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid input", err.Error()))
	}

	// Find user
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Invalid credentials", "Email or password is incorrect"))
	}

	// Verify password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Invalid credentials", "Email or password is incorrect"))
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Error generating token", err.Error()))
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Login successful", "token": token})
}
