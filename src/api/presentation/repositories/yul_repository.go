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
	return r.GetFlights(consts.FlightTypeArrival)
}

func (r *YulRepository) GetDepartures() []models.Flight {
	return r.GetFlights(consts.FlightTypeDeparture)
}

func (r *YulRepository) GetFlights(type_ consts.FlightType) []models.Flight {
	var (
		data []string
		err  error
	)
	if type_ == consts.FlightTypeArrival {
		data, err = r.ScrapperClient.GetArrivals()
	} else {
		data, err = r.ScrapperClient.GetDepartures()
	}
	if err != nil {
		// TODO Handle error appropriately
		return nil
	}
	flights := make([]models.Flight, 0, len(data))
	for _, flight := range data {
		flights = append(flights, clients.MapYulFlightData(flight, consts.FlightTypeDeparture))
	}
	return flights
}
