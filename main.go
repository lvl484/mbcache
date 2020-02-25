package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbhost = "localhost"
	dbport = 5432
	dbuser = "cacheuser"
	dbpass = "pgpassword"
	dbname = "postgres"
)

func main() {
	c.cache = make(map[string]CacheValue)
	db := DBInit()
	defer db.Close()
	defer close(queueCache)
	getFromDB(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	go delTracker(db)
	go queueTracker(db)

	err := http.ListenAndServe(":"+port, newRouter())
	if err != nil {
		log.Println(err)
	}
}
