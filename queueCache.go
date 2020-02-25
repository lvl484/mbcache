package main

import (
	"database/sql"
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
			createInDB(db, cacheDB.data)
		case opUpdate:
			updateInDB(db, cacheDB.data)
		case opDelete:
			delFromBD(db, cacheDB.data.Key)
		}

	}
}
