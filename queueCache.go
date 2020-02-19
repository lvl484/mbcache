package main

const capOfQueue = 1000

var queueCache chan queueData = make(chan queueData, capOfQueue)

func toQueue(data JsonBodyValue, operation int) {
	var c queueData = queueData{operation, data}
	queueCache <- c
}
func queueTracker() {

}
