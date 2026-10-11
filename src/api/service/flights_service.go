package service

import (
	"errors"
	"fmt"
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
	nowStr := time.Now().Format(consts.DATETIME_LAYOUT)
	now, _ := time.Parse(consts.DATETIME_LAYOUT, nowStr)
	flights := fs.GetDepartures()
	// Order latest to earliest
	sort.Slice(flights, func(i, j int) bool {
		iTime, err1 := time.Parse(consts.DATETIME_LAYOUT, flights[i].DepartureTime)
		jTime, err2 := time.Parse(consts.DATETIME_LAYOUT, flights[j].DepartureTime)
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing time:", err1, err2)
			// TODO Handle error appropriately returning API Error
			return false
		}
		return iTime.After(jTime)
	})

	// TODO There must be a faster way to access the closest departure time, but for now this will do
	for _, flight := range flights {
		departureTime, err := time.Parse(consts.DATETIME_LAYOUT, flight.DepartureTime)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			// TODO Handle error appropriately returning API Error
			return nil, errors.New("error parsing departure time")
		}
		if departureTime.Before(now) {
			return flight, nil
		}
	}
	return nil, errors.New("no upcoming departures found")
}
