package main

func toQueue(data JsonBodyValue, operation int) {
	var c queueData = queueData{operation, data}
	queueCache <- c
}
