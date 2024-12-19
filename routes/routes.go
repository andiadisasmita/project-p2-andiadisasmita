package routes

import (
	"os"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/controllers"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupRoutes initializes all routes
func SetupRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom error handler
	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler

	// Public routes
	e.POST("/users/register", controllers.RegisterUser)
	e.POST("/users/login", controllers.LoginUser)

	// Protected routes
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	// Protected routes for rentals, payments, reviews, etc.
	r.GET("/boardgames", controllers.GetBoardgames)
	r.GET("/boardgames/:id", controllers.GetBoardgameByID)
	r.GET("/rentals", controllers.GetUserRentals)
	r.POST("/rentals", controllers.CreateRental)
	r.PUT("/rentals/:id", controllers.UpdateRental)
	r.GET("/rentals/history", controllers.GetRentalHistory)
	r.GET("/payments", controllers.GetPayments)
	r.POST("/payments", controllers.CreatePayment)
	r.GET("/reviews/:boardgame_id", controllers.GetReviews)
	r.POST("/reviews", controllers.CreateReview)
}
