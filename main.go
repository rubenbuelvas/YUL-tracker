package main

import (
	"github.com/rubenbuelvas/yul-tracker/src/api/config"
	"github.com/rubenbuelvas/yul-tracker/src/api/infrastructure/clients"
	"github.com/rubenbuelvas/yul-tracker/src/api/presentation/repositories"
	"github.com/rubenbuelvas/yul-tracker/src/api/presentation/routes"
	"github.com/rubenbuelvas/yul-tracker/src/api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load configuration
	config.LoadConfig()

	// Initialize repositories
	yulRepository := repositories.NewYulRepository(clients.NewYulScrapper())

	// Initialize services
	flightsService := service.NewFlightsService(yulRepository)

	// Initialize API routes
	routes.SetupRoutes(r, flightsService)

	// Run the server
	r.Run()
}
