package main

import (
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", PingHandler)
	http.ListenAndServe(":8080", nil)
}
