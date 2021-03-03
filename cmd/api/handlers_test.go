package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)
	app := Application{
		clock:    staticClock{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		infoLog:  logger,
		errorLog: logger,
	}
	app.now(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("want %s; got %s", http.StatusText(http.StatusOK), http.StatusText(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if content := string(body); content != "2021-01-01T00:00:00Z" {
		t.Errorf("want body to equal %q; got %q", "OK", content)
	}
}
