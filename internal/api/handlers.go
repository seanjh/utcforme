package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type appClock interface {
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

// App contains application-level configuration and context.
type App struct {
	// Application clock.
	clock appClock
	// Info-level logger.
	logInfo *log.Logger
	// Error-level logger.
	logErr *log.Logger
}

// DefaultApp returns a new App.
func DefaultApp(info, err *log.Logger) App {
	return App{
		clock:   defaultClock{},
		logInfo: info,
		logErr:  err,
	}
}

func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	app.logErr.Printf("cannot find path: %s", r.URL.Path)
	http.NotFound(w, r)
}

func (app *App) Now(w http.ResponseWriter, r *http.Request) {
	t := app.clock.Now().UTC()

	var loc *time.Location
	tz := r.URL.Query().Get("zone")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		app.logErr.Println(err.Error())
		http.Error(w, fmt.Sprintf("invalid zone: %s", tz), http.StatusBadRequest)
		return
	}

	lt := t.In(loc)
	w.Write([]byte(lt.Format(time.RFC3339)))
}
