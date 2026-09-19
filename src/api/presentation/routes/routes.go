package routes

import (
	"github.com/rubenbuelvas/yul-tracker/src/api/presentation/handlers"
	"github.com/rubenbuelvas/yul-tracker/src/api/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, flight_service *service.FlightsService) {
	pingHandler := handlers.NewPingHandler()
	flightsHandler := handlers.NewFlightsHandler(flight_service)

	r.GET("/ping", pingHandler.GetPing)

	flightsGroup := r.Group("/flights")
	flightsGroup.GET("/next-arrival", flightsHandler.GetNextArrival)
}
