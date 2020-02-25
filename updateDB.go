package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func clearDB(db *sql.DB) {
	sqlStatments := `DELETE FROM fastcache WHERE DELTIME <$1;`
	_, err := db.Exec(sqlStatments, time.Now())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("cache added to db")
	}

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
		c.Lock()
		temp.Key = temp2.Key
		temp.Value = temp2.Value
		layout := "2006-01-02 15:04:05-07:00"
		t, _ := time.Parse(layout, temp2.Deltime)
		fmt.Println(temp2.Deltime)
		temp.Deltime = &t
		c.cache[temp.Key] = CacheValue{temp.Value, temp.Deltime}
		fmt.Println(c.cache)
		c.Unlock()
	}

}
