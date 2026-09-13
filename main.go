package main

import (
	"yul-tracker/api"
	"yul-tracker/config"
	"yul-tracker/repository"
	"yul-tracker/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load configuration
	config.LoadConfig()

	// Initialize repositories
	repo := repository.NewRepository()

	// Initialize services
	svc := service.NewService(repo)

	// Initialize API routes
	api.SetupRoutes(r, svc)

	// Run the server
	r.Run()
}
