package routes

import (
	"yul-tracker/api/handlers"
	"yul-tracker/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, svc *service.Service) {
	h := handlers.NewHandler(svc)

	r.GET("/ping", h.Ping)
}
