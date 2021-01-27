package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr string
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags|log.Lmicroseconds|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.LUTC)

	app := Application{
		clock:    defaultClock{},
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("localhost/now", app.now)

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("starting server on %s", cfg.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
