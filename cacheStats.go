package main

type cacheStats struct {
	NumOfRec    int `json:"NumOfRec"`
	NumOfAdd    int `json:"NumOfAdd"`
	NumOfDel    int `json:"NumOfDel"`
	NumOfGet    int `json:"NumOfGet"`
	NumOfUpsert int `json:"NumOfUpsert"`
	NumOfUpdate int `json:"NumOfUpdate"`
}
