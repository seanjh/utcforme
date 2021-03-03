package main

import (
	"flag"
	"log"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/mux"
)

type config struct {
	addr      string
	baseRoute string
}

const defaultAddr = ":3000"
const defaultBaseRoute = "/v1/api"

func main() {
	cfg := new(config)
	flag.StringVar(&cfg.addr, "addr", defaultAddr, "HTTP network address")
	flag.StringVar(&cfg.baseRoute, "route", defaultBaseRoute, "Base API route")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags|log.Lmicroseconds|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.LUTC)

	app := Application{
		clock:    defaultClock{},
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	r := mux.NewRouter().PathPrefix(cfg.baseRoute).Subrouter()
	r.HandleFunc("/", app.index)
	r.HandleFunc("/now", app.now)

	//srv := &http.Server{
	//Addr:     cfg.Addr,
	//ErrorLog: errorLog,
	//Handler:  mux,
	//}

	infoLog.Printf("starting server on %s", cfg.addr)
	gateway.ListenAndServe(cfg.addr, r)
	//errorLog.Fatal(srv.ListenAndServe())
}
