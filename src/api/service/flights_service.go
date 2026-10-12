package service

import (
	"fmt"
	"slices"
	"sort"
	"time"

	"github.com/rubenbuelvas/yul-tracker/src/api/domain/consts"
	"github.com/rubenbuelvas/yul-tracker/src/api/domain/models"
	"github.com/rubenbuelvas/yul-tracker/src/api/presentation/repositories"
)

type FlightsService struct {
	yulRepository *repositories.YulRepository
}

func NewFlightsService(yulRepository *repositories.YulRepository) *FlightsService {
	return &FlightsService{
		yulRepository: yulRepository,
	}
}

func (fs *FlightsService) GetArrivals() []*models.Flight {
	return fs.yulRepository.GetArrivals()
}

func (fs *FlightsService) GetDepartures() []*models.Flight {
	return fs.yulRepository.GetDepartures()
}

func (fs *FlightsService) GetLastDeparture(type_ string) (*models.Flight, error) {
	flights := slices.DeleteFunc(fs.GetDepartures(), func(flight *models.Flight) bool {
		return flight.LocationStatus != consts.LocationStatusFlying
	})
	// Order latest to earliest
	sort.Slice(flights, func(i, j int) bool {
		iTime, err1 := time.Parse(consts.DATETIME_LAYOUT, flights[i].DepartureTime)
		jTime, err2 := time.Parse(consts.DATETIME_LAYOUT, flights[j].DepartureTime)
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing time:", err1, err2)
			panic(fmt.Sprintf("Error parsing time:", err1, err2))
		}
		return iTime.After(jTime)
	})
	return flights[0], nil
}
