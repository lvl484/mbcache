package main

import (
	"database/sql"
	"log"
	"time"
)

//Create new cache in BD
func updateDBCreate(data *JsonBodyValue) {

}

//update cache in DB
func updateDBUpdate(data *JsonBodyValue) {

}

//delete cache from db
func updateDBDelete(delKey string) {

}

//get all data to cache when server starts
func getFromDB(db *sql.DB) {
	rows, err := db.Query(
		`SELECT *
       FROM fastcache`)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var temp JsonBodyValue
		type strToTime struct {
			Key     string
			Value   string
			Deltime string
		}
		var temp2 strToTime
		err = rows.Scan(&temp2.Key, &temp2.Value, &temp2.Deltime)

		if err != nil {
			log.Println(err)
			continue
		}

		if err = rows.Err(); err != nil {
			log.Println(err)
			continue
		}

		temp.Key = temp2.Key
		temp.Value = temp2.Value
		layout := "2006-01-02T15:04:05-0700"
		t, _ := time.Parse(layout, temp2.Deltime)
		temp.Deltime = &t
		c.cache[temp.Key] = CacheValue{temp.Value, temp.Deltime}

	}

}
