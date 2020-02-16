package main

import "net/http"

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
