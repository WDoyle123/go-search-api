package calculations

import (
	"math"
)

const RadiusOfEarthKM = 6371008.7714 / 1000

type Coordinate struct {
	latitude  float64
	longitude float64
}

func NewCoordinate(lat, lon float64) Coordinate {
	return Coordinate{latitude: lat, longitude: lon}
}

func convertToRadians(d float64) float64 {
	return d * (math.Pi / 180)
}

func roundTo1DP(x float64) float64 {
	return math.Round(x*10) / 10
}

func Haversine(p1, p2 Coordinate) float64 {
	lat1 := convertToRadians(p1.latitude)
	lon1 := convertToRadians(p1.longitude)

	lat2 := convertToRadians(p2.latitude)
	lon2 := convertToRadians(p2.longitude)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Pow(math.Sin(dLat/2), 2) + (math.Cos(lat1)*math.Cos(lat2))*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := RadiusOfEarthKM * c

	return roundTo1DP(d)
}
