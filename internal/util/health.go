package util

import (
	"net/http"
	"sync"

	"github.com/braintree/manners"
)

var (
	healthzStatus   = http.StatusOK
	readinessStatus = http.StatusOK
	mu              sync.RWMutex
)

func RunHealthServer(healthAddr *string, errChan chan error) {
	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", HealthzHandler)
	hmux.HandleFunc("/readiness", ReadinessHandler)
	hmux.HandleFunc("/healthz/status", HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", ReadinessStatusHandler)
	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hmux.ServeHTTP(w, r)
	})

	errChan <- healthServer.ListenAndServe()
}

func HealthzStatus() int {
	mu.RLock()
	defer mu.RUnlock()
	return healthzStatus
}

func ReadinessStatus() int {
	mu.RLock()
	defer mu.RUnlock()
	return readinessStatus
}

func SetHealthzStatus(status int) {
	mu.Lock()
	healthzStatus = status
	mu.Unlock()
}

func SetReadinessStatus(status int) {
	mu.Lock()
	readinessStatus = status
	mu.Unlock()
}

// HealthzHandler responds to health check requests.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(HealthzStatus())
}

// ReadinessHandler responds to readiness check requests.
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(ReadinessStatus())
}

func ReadinessStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch ReadinessStatus() {
	case http.StatusOK:
		SetReadinessStatus(http.StatusServiceUnavailable)
	case http.StatusServiceUnavailable:
		SetReadinessStatus(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}

func HealthzStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch HealthzStatus() {
	case http.StatusOK:
		SetHealthzStatus(http.StatusServiceUnavailable)
	case http.StatusServiceUnavailable:
		SetHealthzStatus(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}
