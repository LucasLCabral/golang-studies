package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close() // garante que esse comando vai ser executado por ultimo.

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}