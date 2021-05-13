package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	ts := httptest.NewServer(Router())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/vmstat")

    if err != nil {
            t.Fatalf("Http GET faild %s", err)
            return
    }

	//Status Code Check
    if res.StatusCode != 200 {
            t.Error("Status code error")
            return
    }
}