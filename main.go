package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func now(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UTC()

	var loc *time.Location
	tz := r.URL.Query().Get("zone")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid zone: %s", tz), http.StatusBadRequest)
		return
	}

	lt := t.In(loc)
	w.Write([]byte(lt.Format(time.RFC3339)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("localhost/now", now)

	log.Println("starting server on :4000")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
