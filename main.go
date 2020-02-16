// this is the first example realization of my
// where are declareted main functions I need adn realixation of them
// delTracker needs additiolan testing because of comparing time( we can compare time.Time with the same, but we have to test string(Svalue.Deltime))
// and time.Time because in the Json bodi it's string or try to change tope in Json body to time.Time.
// where is no function which deletes 1time per day all out of date cache.
package main

import (
	"log"
	"net/http"
	"os"
)

var cache map[string]CacheValue

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	go delTracker()

	err := http.ListenAndServe(":"+port, newRouter())
	if err != nil {
		log.Println(err)
	}
}
