package clients

import (
	"regexp"
	"strings"

	"github.com/rubenbuelvas/yul-tracker/src/api/domain/consts"
	"github.com/rubenbuelvas/yul-tracker/src/api/domain/models"
)

var (
	time24Regex = regexp.MustCompile(`^([01][0-9]|2[0-3]):[0-5][0-9]$`)
	// TODO This is plagued by assumptions
	arrivalLocationStatusMap = map[string]consts.LocationStatus{
		"Scheduled": consts.LocationStatusScheduled,
		"Cancelled": consts.LocationStatusCancelled,
		"Arrived":   consts.LocationStatusLanded,
		// TODO On time could be scheduled or flying, but for now we will assume that it is flying
		"On time": consts.LocationStatusFlying,
		"Delayed": consts.LocationStatusFlying,
		"Early":   consts.LocationStatusFlying,
	}
	departureLocationStatusMap = map[string]consts.LocationStatus{
		"Departed":  consts.LocationStatusFlying,
		"Cancelled": consts.LocationStatusCancelled,
		"On Time":   consts.LocationStatusScheduled,
		"Delayed":   consts.LocationStatusScheduled,
	}
)

func MapYulFlightData(data string, type_ consts.FlightType) models.Flight {
	// TODO: IDK if logic to map the flight data to the Flight struct should be implemented here
	flight := models.Flight{}
	// Example data:
	// 17:16\n17:30\nUA1925 United AirlinesChicago Ohare (ORD)\nCodeshare flight(s):\nNZ9856\nNH7804\nEK6670\nCM2819\nAC3335\nArrived\nGate 80
	dataList := strings.Split(data, "\n")

	if time24Regex.MatchString(dataList[1]) {
		if compareTimeStrings(dataList[0], dataList[1]) > 0 {
			flight.TimeStatus = consts.TimeStatusDelayed
		} else if compareTimeStrings(dataList[0], dataList[1]) < 0 {
			flight.TimeStatus = consts.TimeStatusEarly
		} else {
			// This should never happen, but just in case, we will set the time status to ON_TIME
			flight.TimeStatus = consts.TimeStatusOnTime
		}
		dataList = append(dataList[:1], dataList[2:]...)
	} else {
		flight.TimeStatus = consts.TimeStatusOnTime
	}

	if type_ == consts.FlightTypeArrival {
		flight.Type = consts.FlightTypeArrival
		flight.ArrivalTime = dataList[0]
		flight.DepartureIATAAirportCode = strings.Split(dataList[1], "(")[1][:3]
		flight.LocationStatus = arrivalLocationStatusMap[dataList[len(dataList)-2]]
	} else {
		flight.Type = consts.FlightTypeDeparture
		flight.DepartureTime = dataList[0]
		flight.ArrivalIATAAirportCode = strings.Split(dataList[1], "(")[1][:3]
		flight.LocationStatus = departureLocationStatusMap[dataList[len(dataList)-2]]
	}
	flight.FlightNumber = strings.Split(dataList[1], " ")[0]

	return flight
}

func compareTimeStrings(time1, time2 string) int {
	// Compare time strings in "HH:MM" format
	if time1 < time2 {
		return -1
	} else if time1 > time2 {
		return 1
	}
	return 0
}
