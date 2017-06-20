package main

import (
	"encoding/json"
	"net/http"

	health "github.com/peteclark-io/health-api"
)

func Health() func(w http.ResponseWriter, r *http.Request) {
	agg := health.Aggregator("users-api", health.Ping())

	return func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(agg())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}
}

func GTG() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
