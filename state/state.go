package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var state float32 = 0.0

func handleState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodOptions && r.Method != http.MethodGet {
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodPost {
		var val float32
		err := json.NewDecoder(r.Body).Decode(&val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Setting state to %f\n", val)
		state = val
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(state)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/state", handleState)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
