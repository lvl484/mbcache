package main

var queueCache chan queueData = make(chan queueData, 1000)

func toQueue(data JsonBodyValue, operation int) {
	var c queueData = queueData{operation, data}
	queueCache <- c
}
func queueTracker() {

}
