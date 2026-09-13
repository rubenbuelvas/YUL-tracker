package main

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	r := gin.Default()

	// Load configuration
	config.LoadConfig()

	// Initialize repositories
	repo := repository.NewRepository()

	// Initialize services
	svc := service.NewService(repo)

	// Initialize API routes
	api.SetupRoutes(r, svc)

	// Perform a request to the /ping endpoint
	w := performRequest(r, "GET", "/ping")

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func performRequest(r *gin.Engine, method, path string) *gin.ResponseWriter {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
