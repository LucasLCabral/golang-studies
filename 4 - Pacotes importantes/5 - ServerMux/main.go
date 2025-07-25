package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler) // Handle the root path with HomeHandler
	mux.Handle("/blog", blog{title: "ZiziBlog"})
	go http.ListenAndServe(":8080", mux) // run in a goroutine

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from mux2!"))
	})
	mux2.Handle("/blog", blog{title: "ZiziBlog in the mux2"})
	http.ListenAndServe(":8081", mux2)   // main goroutine blocks here
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
}

type blog struct{
	title string

}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(b.title + "\n"))
}