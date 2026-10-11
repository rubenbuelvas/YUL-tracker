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

func (h *FlightsHandler) GetArrivals(c *gin.Context) {
	c.JSON(200, h.flightsService.GetArrivals())
}

func (h *FlightsHandler) GetDepartures(c *gin.Context) {
	c.JSON(200, h.flightsService.GetDepartures())
}

func (h *FlightsHandler) GetLastDeparture(c *gin.Context) {
	flight, err := h.flightsService.GetLastDeparture("departure")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, flight)
}
