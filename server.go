package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const (
	key        = "Key"
	defDelTime = time.Hour
	opCreate   = 0
	opDelete   = 1
	opUpdate   = 2
	capOfQueue = 1000
)

type safeCache struct {
	sync.Mutex
	cache map[string]CacheValue
}

var c safeCache
var stats cacheStats

// addCache add cache to RAM and db
func addCache(w http.ResponseWriter, r *http.Request) {

	var reqcache JsonBodyValue
	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if reqcache.Deltime == nil {
		tempt := time.Now().Add(defDelTime)
		reqcache.Deltime = &tempt
	}
	c.Lock()
	if _, ok := c.cache[reqcache.Key]; ok {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Such key is already exist"))
		c.Unlock()
		return
	}
	c.cache[reqcache.Key] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	toQueueCreate(reqcache)
	stats.NumOfAdd++
	c.Unlock()
	w.WriteHeader(http.StatusOK)

}

func upsertCache(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reqcache JsonBodyValue
	err := json.NewDecoder(r.Body).Decode(&reqcache)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	reqcache.Key = params[key]
	if reqcache.Deltime == nil {
		tempt := time.Now().Add(defDelTime)
		reqcache.Deltime = &tempt
	}
	c.Lock()
	if _, ok := c.cache[reqcache.Key]; ok {
		toQueueUpdate(reqcache)
	} else {
		toQueueCreate(reqcache)
	}
	c.cache[reqcache.Key] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	stats.NumOfUpsert++
	c.Unlock()
	w.WriteHeader(http.StatusOK)

}

// getOneCache gets cache my Key in URL as a param
func getOneCache(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	c.Lock()
	if _, ok := c.cache[params[key]]; !ok {
		w.WriteHeader(http.StatusNotFound)
		c.Unlock()
		return

	}
	err := json.NewEncoder(w).Encode(c.cache[params[key]])
	stats.NumOfGet++
	c.Unlock()
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// deleteCache deletes cache from RAM(map)
func deleteCache(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	c.Lock()
	if _, ok := c.cache[params[key]]; !ok {
		c.Unlock()
		w.WriteHeader(http.StatusNotFound)
		return
	}
	toQueueDelete(params[key])
	delete(c.cache, params[key])
	stats.NumOfDel++
	c.Unlock()
	w.WriteHeader(http.StatusOK)

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
	}
	if _, ok := c.cache[reqcache.Key]; !ok {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Such key is not exist"))
		c.Unlock()
		return
	}
	if reqcache.Deltime == nil {
		tempt := time.Now().Add(defDelTime)
		reqcache.Deltime = &tempt
	}
	c.Lock()
	c.cache[params[key]] = CacheValue{
		reqcache.Value,
		reqcache.Deltime,
	}
	toQueueUpdate(reqcache)
	stats.NumOfUpdate++
	c.Unlock()
	w.WriteHeader(http.StatusOK)

}

//getStats gives info about "Stats of cache, number of records, memory consumption..
func getStats(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(stats)
	if err != nil {
		log.Println(err)
	}

}

// delTracker deletes cache from RAM every minute
func delTracker() {

	for {
		c.Lock()
		for k, v := range c.cache {
			if v.Deltime.Before(time.Now()) {
				delete(c.cache, k)
			}
		}
		c.Unlock()
		time.Sleep(time.Minute)
	}
}
