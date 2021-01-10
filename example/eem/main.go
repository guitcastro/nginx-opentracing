package main

import (
	"go.elastic.co/apm/module/apmhttp"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("hello world"))
}

func main() {
	tracedHandler := apmhttp.Wrap(http.HandlerFunc(handler))
	http.Handle("/hello", tracedHandler)
	http.ListenAndServe(":9001", nil)
}
