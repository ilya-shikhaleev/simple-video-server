package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", list).Methods(http.MethodGet)
	s.HandleFunc("/video/d290f1ee-6c54-4b01-90e6-d701748f0851", video).Methods(http.MethodGet)
	return r
}
