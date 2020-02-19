package main

var queueCache chan queueData = make(chan queueData, capOfQueue)

func toQueue(data JsonBodyValue, operation int) {
	var c queueData = queueData{operation, data}
	queueCache <- c
}
func queueTracker() {
	//add chech chal is empty and op to db with cases opDelete,opUpdate,opCreate
}
