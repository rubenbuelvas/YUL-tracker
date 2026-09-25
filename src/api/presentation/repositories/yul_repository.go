package repositories

import (
	"github.com/rubenbuelvas/yul-tracker/src/api/domain/consts"
	"github.com/rubenbuelvas/yul-tracker/src/api/domain/models"
	"github.com/rubenbuelvas/yul-tracker/src/api/infrastructure/clients"
)

type YulRepository struct {
	ScrapperClient *clients.YulScrapper
}

func NewYulRepository(scrapperClient *clients.YulScrapper) *YulRepository {
	return &YulRepository{
		ScrapperClient: scrapperClient,
	}
}

func (r *YulRepository) GetArrivals() []models.Flight {
	data, err := r.ScrapperClient.GetArrivals()
	if err != nil {
		// Handle error appropriately
		return nil
	}
	flights := make([]models.Flight, 0, len(data))
	for _, flight := range data {
		flights = append(flights, clients.MapYulFlightData(flight, consts.FlightTypeArrival))
	}
	return flights
}
