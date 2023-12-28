package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Operands struct {
	OperandOne float32 `json:"operandOne"`
	OperandTwo float32 `json:"operandTwo"`
}

func subtract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var op Operands
	err := json.NewDecoder(r.Body).Decode(&op)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Subtracting %f from %f\n", op.OperandTwo, op.OperandOne)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(op.OperandOne - op.OperandTwo)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/subtract", subtract)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
