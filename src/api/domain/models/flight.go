package models

import "github.com/rubenbuelvas/yul-tracker/src/api/domain/consts"

type Flight struct {
	FlightNumber             string                `json:"flight_number"`
	DepartureIATAAirportCode string                `json:"departure_iata_airport_code"`
	ArrivalIATAAirportCode   string                `json:"arrival_iata_airport_code"`
	DepartureTime            string                `json:"departure_time"`
	ArrivalTime              string                `json:"arrival_time"`
	LocationStatus           consts.LocationStatus `json:"location_status"`
	TimeStatus               consts.TimeStatus     `json:"time_status"`
}
