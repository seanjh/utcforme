package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/seanjh/utcforme/internal/api"
)

type config struct {
	addr      string
	baseRoute string
}

const defaultAddr = ":3000"
const defaultBaseRoute = "/v1/api"

func boom(w http.ResponseWriter, r *http.Request) {
	panic("oh heck")
}

func main() {
	cfg := new(config)
	flag.StringVar(&cfg.addr, "addr", defaultAddr, "HTTP network address")
	flag.StringVar(&cfg.baseRoute, "route", defaultBaseRoute, "Base API route")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags|log.Lmicroseconds|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.LUTC)

	app := api.DefaultApp(infoLog, errorLog)

	r := mux.NewRouter().PathPrefix(cfg.baseRoute).Subrouter()
	r.HandleFunc("/", app.Index)
	r.HandleFunc("/now", app.Now)
	r.HandleFunc("/boom", boom)

	infoLog.Printf("starting server on %s", cfg.addr)
	gateway.ListenAndServe(cfg.addr, handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
	)(r))
}
