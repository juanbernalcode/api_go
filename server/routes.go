package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	// Route == "/"
	http.HandleFunc("/", index)

	// Route == "/countries"
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET
		case http.MethodGet:
			getCountries(w, r)
		//	POST
		case http.MethodPost:
			addCountrie(w, r)
		// others methods
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Method not allowed")
			return
		}
	})

}
