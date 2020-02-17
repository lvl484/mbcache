package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const key = "Key"

type safeCache struct {
	sync.Mutex
	cache map[string]CacheValue
}

var c safeCache

var stats cacheStats

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
	c.Lock()
	if _, ok := c.cache[reqcache.Key]; ok {
		w.Write([]byte("Such key is already exist"))
		c.Unlock()
		return
	}
	c.cache[reqcache.Key] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	c.Unlock()
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
	c.Lock()
	c.cache[reqcache.Key] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	c.Unlock()
	w.WriteHeader(http.StatusOK)
}

// getOneCache gets cache my Key in URL as a param
func getOneCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//togo adding to queue
	if _, ok := c.cache[params[key]]; ok {
		err := json.NewEncoder(w).Encode(c.cache[params[key]])
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
	//todo add func in queue
	if _, ok := c.cache[params[key]]; ok {
		delete(c.cache, params[key])
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
		c.cache[params[key]] = CacheValue{
			reqcache.Value,
			reqcache.Deltime,
		}
	}
}

//getStats gives info about "Stats of cache, number of records, memory consumption..
func getStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(stats)
	if err != nil {
		log.Println(err)
	}

}

// delTracker deletes cache from RAM every minute
func delTracker() {
	//todo add req to bd (queue)
	for {
		for k, v := range c.cache {
			if v.Deltime.Before(time.Now()) {
				delete(c.cache, k)
			}
		}
		time.Sleep(time.Minute)
	}
}

//queueTracker works with db and queue
func queueTracker() {

}
