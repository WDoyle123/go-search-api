package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"go-search-api/calculations"
	"go-search-api/internal/database"
)

func health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func root(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"service": "go-search-api"})
}

func search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	eventIDStr := q.Get("event_id")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid event_id"})
		return
	}

	radiusStr := q.Get("radius_km")
	radiusKM, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid radius_km"})
		return
	}

	allModes := []string{"walking", "driving", "public_transport"}

	transportModesStr := q.Get("modes")
	var transportModes []string
	if transportModesStr == "" {
		transportModes = allModes
	} else {
		raw := strings.SplitSeq(transportModesStr, ",")
		for tm := range raw {
			tm = strings.TrimSpace(tm)
			if tm != "" {
				transportModes = append(transportModes, tm)
			}
		}
		if len(transportModes) == 0 {
			transportModes = allModes
		}
	}

	var event database.Event
	found := false
	for _, e := range database.Events {
		if e.ID == eventID {
			event = e
			found = true
			break
		}
	}

	if !found {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "event not found"})
		return
	}

	result := []SearchResponse{}

	eventCoordinate := calculations.NewCoordinate(event.Latitude, event.Longitude)
	for _, h := range database.Hotels {
		hotelCoordinate := calculations.NewCoordinate(h.Latitude, h.Longitude)
		d := calculations.Haversine(eventCoordinate, hotelCoordinate)
		if d >= radiusKM {
			continue
		}
		travelTimes := make(map[string]float64)

		for _, tm := range transportModes {
			t, err := calculations.TravelTime(tm, d)
			if err != nil {
				log.Println(err)
				continue
			}
			travelTimes[tm] = t
		}
		result = append(result, SearchResponse{
			HotelID:         h.ID,
			Latitude:        h.Latitude,
			Longitude:       h.Longitude,
			DistanceKm:      d,
			EstimatedTravel: travelTimes,
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{"search": result})
}

type SearchResponse struct {
	HotelID         int                `json:"hotel_id"`
	Latitude        float64            `json:"latitude"`
	Longitude       float64            `json:"longitude"`
	DistanceKm      float64            `json:"distance_km"`
	EstimatedTravel map[string]float64 `json:"estimated_travel"`
}
