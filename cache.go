package main

import "time"

type JsonBodyValue struct {
	Key     string     `json:"Key"`
	Value   string     `json:"Value"`
	Deltime *time.Time `json:"Deltime"`
}

type CacheValue struct {
	Value   string
	Deltime *time.Time
}
type middleDataFromDB struct {
	Key     string
	Value   string
	Deltime string
}
