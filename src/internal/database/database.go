package database

type Event struct {
	ID        int
	Latitude  float64
	Longitude float64
}

type Hotel struct {
	ID        int
	Latitude  float64
	Longitude float64
}

var Events = []Event{
	{ID: 1, Latitude: 51.5007, Longitude: -0.1246},
}

var Hotels = []Hotel{
	{ID: 101, Latitude: 51.5033, Longitude: -0.1195},
	{ID: 102, Latitude: 51.5094, Longitude: -0.1183},
	{ID: 103, Latitude: 51.4952, Longitude: -0.1469},
	{ID: 104, Latitude: 51.5155, Longitude: -0.0720},
	{ID: 105, Latitude: 51.4700, Longitude: -0.4543},
}
