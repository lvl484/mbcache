package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// addCache add cache to RAM and db
func addCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqcache JsonBodyValue
	// todo add addint to queue
	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if _, ok := cache[reqcache.Key]; ok {
		w.Write([]byte("Such key is already exist"))
		return
	}
	cache[reqcache.Key] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	w.WriteHeader(http.StatusOK)
}

func upsertCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqcache JsonBodyValue
	// todo add addint to queue
	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	cache[reqcache.Key] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	w.WriteHeader(http.StatusOK)
}

// getOneCache gets cache my Key in URL as a param
func getOneCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//togo dding to queue
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

	var reqcache JsonBodyValue

	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		cache[params["Key"]] = CacheValue{
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
	stats.numOfRec = len(cache)
}

// delTracker deletes cache from RAM every minute
func delTracker() {
	for {
		for k, v := range cache {
			if v.Deltime.Before(time.Now()) {
				delete(cache, k)
			}
		}
		time.Sleep(time.Minute)
	}
}
