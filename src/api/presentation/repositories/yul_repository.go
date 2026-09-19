package repositories

import (
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

func (r *YulRepository) GetNextArrival() string {
	return r.ScrapperClient.GetNextArrival()
}
