package main

import (
	"database/sql"
	"log"
)

var queueCache chan queueData = make(chan queueData, capOfQueue)

func toQueueCreate(data JsonBodyValue) {
	var c queueData = queueData{opCreate, data}
	queueCache <- c
}
func toQueueUpdate(data JsonBodyValue) {
	var c queueData = queueData{opUpdate, data}
	queueCache <- c
}
func toQueueDelete(qkey string) {
	var data JsonBodyValue = JsonBodyValue{qkey, "", nil}
	var c queueData = queueData{opDelete, data}
	queueCache <- c
}
func queueTracker(db *sql.DB) {

	for {
		cacheDB := <-queueCache
		switch cacheDB.operaion {
		case opCreate:
			sqlStatments := `INSERT INTO fastcache (CKEY, VALUE, DELTIME)
		VALUES ($1, $2, $3)`
			_, err := db.Exec(sqlStatments, cacheDB.data.Key, cacheDB.data.Value, cacheDB.data.Deltime)
			if err != nil {
				log.Println(err)
			}
		case opUpdate:
			sqlStatments := `UPDATE fastcache SET VALUE=$2,DELTIME=$3
		WHERE CKEY=$1;`
			_, err := db.Exec(sqlStatments, cacheDB.data.Key, cacheDB.data.Value, cacheDB.data.Deltime)
			if err != nil {
				log.Println(err)
			}
		case opDelete:
			sqlStatments := `DELETE FROM fastcache WHERE CKEY=$1;`
			_, err := db.Exec(sqlStatments, cacheDB.data.Key)
			if err != nil {
				log.Println(err)
			}

		}

	}
}
