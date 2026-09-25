package service

import (
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

func (fs *FlightsService) GetArrivals() []models.Flight {
	return fs.yulRepository.GetArrivals()
}
