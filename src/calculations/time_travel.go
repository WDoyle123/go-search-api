package calculations

import (
	"errors"
	"fmt"
	"time"
)

const (
	MinimumDistance = 0
	MaximumDistance = 250
)

var (
	ErrDistanceTooSmall     = errors.New("distance must be greater than minimum")
	ErrDistanceTooLarge     = errors.New("distance must be less than maximum")
	ErrInvalidTransportMode = errors.New("invalid mode of transport chosen")
)

type TransportMode struct {
	Name         string
	SpeedKmH     float64
	OverheadTime time.Duration
}

var (
	Walking = TransportMode{
		Name:         "walking",
		SpeedKmH:     5,
		OverheadTime: 0,
	}

	Driving = TransportMode{
		Name:         "driving",
		SpeedKmH:     30,
		OverheadTime: 0,
	}

	PublicTransport = TransportMode{
		Name:         "public_transport",
		SpeedKmH:     25,
		OverheadTime: 8 * time.Minute,
	}
)

func ConvertHoursToSeconds(t float64) float64 {
	return t * 3600
}

func ConvertMinutesToSeconds(t float64) float64 {
	return t * 60
}

func TravelTime(tm string, d float64) (float64, error) {
	if d < MinimumDistance {
		return 0, fmt.Errorf("distance must be greater than %d: %w", MinimumDistance, ErrDistanceTooSmall)
	}

	if d > MaximumDistance {
		return 0, fmt.Errorf("distance must be less than %d: %w", MaximumDistance, ErrDistanceTooLarge)
	}

	if tm == "walking" {
		return ConvertHoursToSeconds(d / 5), nil
	}

	if tm == "driving" {
		return ConvertHoursToSeconds(d / 30), nil
	}

	if tm == "publicTransport" {
		return ConvertHoursToSeconds(d/25) + ConvertMinutesToSeconds(8), nil
	}

	return 0, errors.New("invalid transport mode")
}
