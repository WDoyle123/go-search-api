package server

import (
	"log"
	"net/http"
	"time"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sw, r)
		ip := getClientIP(r)

		log.Printf("%s %s status=%d ip=%s duration=%s", r.Method, r.URL.Path, sw.status, ip, time.Since(start))
	})
}

