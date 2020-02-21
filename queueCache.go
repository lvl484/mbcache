package main

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
	var data JsonBodyValue = JsonBodyValue{qkey, nil, nil}
	var c queueData = queueData{opDelete, data}
	queueCache <- c
}
func queueTracker() {
	//add chech chal is empty and op to db with cases opDelete,opUpdate,opCreate
}
