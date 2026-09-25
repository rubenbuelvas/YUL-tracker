package consts

type LocationStatus string

const (
	LocationStatusScheduled LocationStatus = "SCHEDULED"
	LocationStatusCancelled LocationStatus = "CANCELLED"
	LocationStatusFlying    LocationStatus = "FLYING"
	LocationStatusLanded    LocationStatus = "LANDED"
)

type TimeStatus string

const (
	TimeStatusOnTime  TimeStatus = "ON_TIME"
	TimeStatusEarly   TimeStatus = "EARLY"
	TimeStatusDelayed TimeStatus = "DELAYED"
)

type FlightType string

const (
	FlightTypeArrival   FlightType = "ARRIVAL"
	FlightTypeDeparture FlightType = "DEPARTURE"
)
