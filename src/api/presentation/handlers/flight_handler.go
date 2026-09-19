package handlers

import (
	"github.com/rubenbuelvas/yul-tracker/src/api/service"

	"github.com/gin-gonic/gin"
)

type FlightsHandler struct {
	flightsService *service.FlightsService
}

func NewFlightsHandler(flightsService *service.FlightsService) *FlightsHandler {
	return &FlightsHandler{flightsService: flightsService}
}

func (h *FlightsHandler) GetNextFlight(c *gin.Context) {
	flight := h.flightsService.GetLanding()
	c.JSON(200, flight)
}
