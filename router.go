package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware)
	router.HandleFunc("/api/cache", addCache).Methods(http.MethodPost)
	router.HandleFunc("/api/cache/{Key}", getOneCache).Methods(http.MethodGet)
	router.HandleFunc("/api/cache/{Key}", deleteCache).Methods(http.MethodDelete)
	router.HandleFunc("/api/cache/u/{Key}", updateCache).Methods(http.MethodPost)
	router.HandleFunc("/api/cache/us/{Key}", upsertCache).Methods(http.MethodPost)
	router.HandleFunc("/api/stats", getStats).Methods(http.MethodGet)
	return router
}
