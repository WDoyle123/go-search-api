package calculations

import "time"

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
