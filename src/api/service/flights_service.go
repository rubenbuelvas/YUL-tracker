package service

import (
	"github.com/rubenbuelvas/yul-tracker/src/api/domain/consts"
	"github.com/rubenbuelvas/yul-tracker/src/api/domain/models"
)

type FlightsService struct {
	// Use repository
}

func NewFlightsService() *FlightsService {
	return &FlightsService{}
}

func (fs *FlightsService) GetLanding() models.Flight {
	return models.Flight{
		FlightNumber:             "AC123",
		DepartureIATAAirportCode: "YUL",
		ArrivalIATAAirportCode:   "JFK",
		DepartureTime:            "2023-06-01T10:00:00Z",
		ArrivalTime:              "2023-06-01T12:00:00Z",
		LocationStatus:           consts.LocationStatusScheduled,
		TimeStatus:               consts.TimeStatusOnTime,
	}
}
