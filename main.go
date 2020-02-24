package main

import (
	"database/sql"
	"fmt"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport,
		dbuser, dbpass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	getFromDB(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	defer close(queueCache)

	go delTracker()
	go queueTracker(db)

	err = http.ListenAndServe(":"+port, newRouter())
	if err != nil {
		log.Println(err)
	}
}
