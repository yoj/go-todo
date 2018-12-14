package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index Router
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex is TODO Top
func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // <- Added
	w.WriteHeader(http.StatusOK)                                      // <- Added

	json.NewEncoder(w).Encode(todos)
}

// TodoShow is TODO Top
func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}
