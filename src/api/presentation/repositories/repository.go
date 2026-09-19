package repository

// Import necessary packages
import (
	"your-scrapper-client-package"
)

// Define the Repository struct
type Repository struct {
	ScrapperClient *your-scrapper-client-package.Client
}

// Initialize the Repository with a scrapper client
func NewRepository() *Repository {
	return &Repository{
		ScrapperClient: your-scrapper-client-package.NewClient(),
	}
}

// Implement methods to interact with the scrapper client
func (r *Repository) FetchData() ([]YourDataType, error) {
	data, err := r.ScrapperClient.FetchData()
	if err != nil {
		return nil, err
	}
	return data, nil
}
