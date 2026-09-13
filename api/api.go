package api

import (
	"yul-tracker/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, svc service.Service) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
