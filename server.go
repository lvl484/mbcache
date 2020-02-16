package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// getCache gets all the cache from RAM
func getCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// addCache add cache to RAM and db
func addCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqcache Jvalue
	// Add addint to db
	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		// todo add err response
	} else {
		if _,ok:=cache;ok{

		}else{
		cache[reqcache.Key] = Svalue{
			reqcache.Value,
			reqcache.Deltime,
		}
	
		w.WriteHeader(http.StatusOK)
	}
}
}

// getOneCache gets cache my Key in URL as a param
func getOneCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
//togo 
	if _, ok := cache[params["Key"]]; ok {
		err := json.NewEncoder(w).Encode(cache[params["Key"]])
		if err != nil {
			log.Println(err)
			return
		} 
			w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

// deleteCache deletes cache from RAM(map)
func deleteCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if _, ok := cache[params["Key"]]; ok {
		delete(cache, params["Key"])
		w.WriteHeader(http.StatusOK)
		return
	} 
		w.WriteHeader(http.StatusNotFound)
	
}

// updateCache updates cache by the key
func updateCache(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var reqcache Jvalue

	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		cache[params["Key"]] = Svalue{
			reqcache.Value,
			reqcache.Deltime,
		}
	}
}

//getStats gives info about "Stats of cache, number of records, memory consumption..
func getStats(w http.ResponseWriter, r *http.Request) {
	type cacheStats struct {
		statsOfCache int
		numOfRec     int
		memCons      int
	}
	var stats cacheStats
	stats.numOfRec := len(cache)
}

// delTracker deletes cache from RAM every minute
func delTracker() {
	for {
		for k, v := range cache {
			//do I need to make 1 more struct for time for comparing
			if v.Deltime < time.Now().Format("01-JAN-2006 15:04") {
				delete(cache, k)
			}
		}
		time.Sleep(time.Minute)
	}
}
