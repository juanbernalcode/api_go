package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handler for route  "/" - (GET)
func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
		return
	}
	fmt.Fprintf(w, "Hello there %s", "Visitor")
}

// handler for route "/countries" - (GET)
func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

// handler for route "/countries" - (POST)
func addCountrie(w http.ResponseWriter, r *http.Request) {
	country := &Country{}

	err := json.NewDecoder(r.Body).Decode(country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	countries = append(countries, country)
	fmt.Fprintln(w, "Country was added")
}
