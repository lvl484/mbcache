package main

import (
	"log"
	"net/http"
	"os"
)

var queueCache chan queueData = make(chan queueData, 1000)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	defer close(queueCache)
	go delTracker()
	//go queueTracker()

	//doesn't work with newRouter
	err := http.ListenAndServe(":"+port, newRouter())
	if err != nil {
		log.Println(err)
	}
}
