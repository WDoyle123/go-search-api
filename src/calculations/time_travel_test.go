package calculations

import (
	"errors"
	"testing"
)

func TestTravelTime(t *testing.T) {
	cases := []struct {
		name          string
		want          float64
		transportMode string
		distance      float64
		wantErr       bool
		wantErrIs     error
	}{
		{name: "walking 1", want: 720, transportMode: "walking", distance: 1, wantErr: false},
		{name: "walking 2", want: 1440, transportMode: "walking", distance: 2, wantErr: false},
		{name: "driving 1", want: 120, transportMode: "driving", distance: 1, wantErr: false},
		{name: "driving 2", want: 240, transportMode: "driving", distance: 2, wantErr: false},
		{name: "public transport 1", want: 624, transportMode: "publicTransport", distance: 1, wantErr: false},
		{name: "public transport 2", want: 768, transportMode: "publicTransport", distance: 2, wantErr: false},
		{name: "bogus mode", transportMode: "horse", distance: 1, wantErr: true, wantErrIs: ErrInvalidTransportMode},
		{name: "distance too small", transportMode: "driving", distance: -1, wantErr: true, wantErrIs: ErrDistanceTooSmall},
		{name: "distance too large", transportMode: "driving", distance: 251, wantErr: true, wantErrIs: ErrDistanceTooLarge},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := TravelTime(c.transportMode, c.distance)

			if (err != nil) != c.wantErr {
				t.Fatalf("error: %v, wantErr: %v", err, c.wantErr)
			}

			if c.wantErr && err != nil {
				if c.wantErrIs != nil && !errors.Is(err, c.wantErrIs) {
					t.Fatalf("error %v does not match sentinel %v", err, c.wantErrIs)
				}
				return
			}

			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}
