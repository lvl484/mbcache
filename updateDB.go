package main

import (
	"database/sql"
	"log"
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
		`SELECT CKEY, VALUE, DELTIME
       FROM fastcache`)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var temp JsonBodyValue
		err = rows.Scan(&temp.Key, &temp.Value, &temp.Deltime)

		if err != nil {
			log.Println(err)
			continue
		}

		if err = rows.Err(); err != nil {
			log.Println(err)
			continue
		}
		c.cache[temp.Key] = CacheValue{temp.Value, temp.Deltime}

	}

}
