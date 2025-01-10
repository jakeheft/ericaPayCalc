package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type inputData struct {
	FullPayHours float32 `json:"fullhours"`
	HalfPayHours float32 `json:"halfhours"`
	LateSegments float32 `json:"latesegments"`
}

const (
	fullRate float32 = 26
	halfRate float32 = 13
)

func calculatePay(w http.ResponseWriter, r *http.Request) {
	var fullPayHours, halfPayHours, lateHours float32
	var data inputData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "err decoding body", http.StatusBadRequest)
	}

	fullPayHours = data.FullPayHours
	lateHours = data.LateSegments * 0.25
	fullPayHours += lateHours
	fullPayTotal := fullPayHours * fullRate
	halfPayHours = data.HalfPayHours
	halfPayTotal := halfPayHours * halfRate
	total := fullPayTotal + halfPayTotal

	fmt.Println("FULL PAY:", fullPayTotal)
	fmt.Println("HALF PAY:", halfPayTotal)
	fmt.Println("TOTAL:", total)
	response := struct{ result float32 }{result: total}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Serve static files (including index.html)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API endpoint
	http.HandleFunc("/calculate", calculatePay)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
