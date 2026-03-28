package server

import (
	"net/http"
	"strconv"

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
	radiusStr := q.Get("radius_km")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error":"invalid event_id")
	}

	radiusKM, err := strconv.Atoi(radiusStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error":"invalid radius_km")
	}

	var result []database.Hotel

	eventCoordinate := calculations.NewCoordinate(e.lat, e.lon)
	for _, h := range database.Hotels {
		hotelCoordinate := calculations.NewCoordinate(h.lat, h.lon)
		d := calculations.Haversine(eventCoordinate, hotelCoordinate)
		if d <= radiusKM {
			result = append(result, h)
		}
	}
	writeJSON(w, http.StatusOK, map[string]string{"search": result})
}
