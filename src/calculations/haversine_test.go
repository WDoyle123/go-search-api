package calculations

import (
	"math"
	"strconv"
	"testing"
)

func TestHaversine(t *testing.T) {
	cases := []struct {
		want float64
		a    Coordinate
		b    Coordinate
	}{
		{want: 0, a: Coordinate{latitude: 0, longitude: 0}, b: Coordinate{latitude: 0, longitude: 0}},
		{want: 111.2, a: Coordinate{latitude: 0, longitude: 0}, b: Coordinate{latitude: 1, longitude: 0}},
		{want: 6161.4, a: Coordinate{latitude: 38.898, longitude: -77.037}, b: Coordinate{latitude: 48.858, longitude: 2.294}},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := Haversine(c.a, c.b)

			if c.want != got {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func TestConvertToRadians(t *testing.T) {
	cases := []struct {
		want   float64
		degree float64
	}{
		{want: 0, degree: 0},
		{want: 2 * math.Pi, degree: 360},
		{want: math.Pi, degree: 180},
		{want: 0.5 * math.Pi, degree: 90},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := convertToRadians(c.degree)
			if c.want != got {
				t.Errorf("got %v, want %v", c.want, got)
			}
		})
	}
}
