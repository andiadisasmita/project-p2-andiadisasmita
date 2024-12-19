package routes

import (
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/controllers"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom error handler
	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler

	// Public routes
	e.POST("/users/register", controllers.RegisterUser)
	e.POST("/users/login", controllers.LoginUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Protected routes
	r := e.Group("") // Create a protected route group
	r.Use(utils.CustomJWTMiddleware)

	// Protected routes for rentals, payments, reviews, etc.
	r.GET("/boardgames", controllers.GetBoardgames)
	r.GET("/boardgames/:id", controllers.GetBoardgameByID)
	r.GET("/rentals", controllers.GetUserRentals)
	r.POST("/rentals", controllers.CreateRental)
	r.PUT("/rentals/:id", controllers.UpdateRental)
	r.GET("/rentals/history", controllers.GetRentalHistory)
	r.GET("/reviews/:boardgame_id", controllers.GetReviews)
	r.POST("/reviews", controllers.CreateReview)

	// New payment routes
	r.POST("/payments/invoice", controllers.CreateInvoice)
	r.GET("/payments/status/:invoice_id", controllers.CheckInvoice)
}
