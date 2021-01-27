package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Clock interface {
	Now() time.Time
}

type defaultClock struct{}

func (defaultClock) Now() time.Time { return time.Now() }

type staticClock struct {
	value time.Time
}

func (c staticClock) Now() time.Time {
	return c.value
}

type Application struct {
	clock    Clock
	infoLog  *log.Logger
	errorLog *log.Logger
}

func index(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (app *Application) now(w http.ResponseWriter, r *http.Request) {
	t := app.clock.Now().UTC()

	var loc *time.Location
	tz := r.URL.Query().Get("zone")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, fmt.Sprintf("invalid zone: %s", tz), http.StatusBadRequest)
		return
	}

	lt := t.In(loc)
	w.Write([]byte(lt.Format(time.RFC3339)))
}
