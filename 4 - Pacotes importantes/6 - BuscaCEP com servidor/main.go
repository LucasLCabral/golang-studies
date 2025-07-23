package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Hello, World!"}`))
}