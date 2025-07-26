package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name  string
	Age   int
}

type DogList []Dog

func main () {
	t := template.Must(template.New("index.html").ParseFiles("index.html"))
	err := t.Execute(os.Stdout, DogList{
		{Name: "Luke", Age: 6},
		{Name: "Mel", Age: 15},
		{Name: "Max", Age: 3},
	})

	if err != nil {
		panic(err)
	}
	
}