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

var transportModes = map[string]TransportMode{
	"walking":          Walking,
	"driving":          Driving,
	"public_transport": PublicTransport,
}

func TravelTime(mode string, d float64) (float64, error) {
	if d < MinimumDistance {
		return 0, fmt.Errorf("distance must be greater than %d: %w", MinimumDistance, ErrDistanceTooSmall)
	}

	if d > MaximumDistance {
		return 0, fmt.Errorf("distance must be less than %d: %w", MaximumDistance, ErrDistanceTooLarge)
	}

	tm, ok := transportModes[mode]
	if !ok {
		return 0, ErrInvalidTransportMode
	}

	hours := d / tm.SpeedKmH
	seconds := hours * 3600

	return seconds + tm.OverheadTime.Seconds(), nil
}
