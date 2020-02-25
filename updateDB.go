package main

import (
	"database/sql"
	"log"
	"time"
)

func delFromBD(db *sql.DB, dkey string) {
	sqlStatments := `DELETE FROM fastcache WHERE CKEY=$1;`
	_, err := db.Exec(sqlStatments, dkey)
	if err != nil {
		log.Println(err)
	}
}
func createInDB(db *sql.DB, body JsonBodyValue) {
	sqlStatments := `INSERT INTO fastcache (CKEY, VALUE, DELTIME)
	VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatments, body.Key, body.Value, body.Deltime)
	if err != nil {
		log.Println(err)
	}
}
func updateInDB(db *sql.DB, body JsonBodyValue) {
	sqlStatments := `UPDATE fastcache SET VALUE=$2,DELTIME=$3
		WHERE CKEY=$1;`
	_, err := db.Exec(sqlStatments, body.Key, body.Value, body.Deltime)
	if err != nil {
		log.Println(err)
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
		var temp2 middleDataFromDB

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
		temp.Key, temp.Value = temp2.Key, temp2.Value
		layout := "2006-01-02 15:04:05-07:00"
		t, _ := time.Parse(layout, temp2.Deltime)
		temp.Deltime = &t
		c.cache[temp.Key] = CacheValue{temp.Value, temp.Deltime}
		c.Unlock()
	}

}
