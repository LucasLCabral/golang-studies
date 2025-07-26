package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request started")
	defer log.Println("request ended")
	select {
		case <-time.After(5 * time.Second):
			log.Println("request processed")
			w.Write([]byte("Request processed succesfully"))
		case <-ctx.Done():
			log.Println("request cancelled by client:", ctx.Err())
	}
}