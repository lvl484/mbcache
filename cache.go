package main

import "time"

type JsonBodyValue struct {
	Key     string     `json:"Key"`
	Value   []byte     `json:"Value"`
	Deltime *time.Time `json:"Deltime"`
}

type CacheValue struct {
	Value   []byte
	Deltime *time.Time
}

type queueData struct {
	operaion string
	data     JsonBodyValue
}
