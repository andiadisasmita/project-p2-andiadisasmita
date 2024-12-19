package main

import (
	"log"
	"os"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Boardgame Rental API
// @version 1.0
// @description This is the REST API for the Boardgame Rental Application.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	if err := config.InitializeDatabase(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize Echo instance
	e := echo.New()

	// Set up routes
	routes.SetupRoutes(e)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
