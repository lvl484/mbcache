package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/cache", addCache).Methods(http.MethodPost)
	router.HandleFunc("/api/cache/{key}", getOneCache).Methods(http.MethodGet)
	router.HandleFunc("/api/cache/{key}", deleteCache).Methods(http.MethodDelete)
	router.HandleFunc("/api/cache/{key}", updateCache).Methods(http.MethodPost)
	router.HandleFunc("/api/stats", getStats).Methods(http.MethodGet)
	return router
}
