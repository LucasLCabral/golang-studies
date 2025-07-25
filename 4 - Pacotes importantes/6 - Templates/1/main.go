package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name  string
	Age   int
}


func main () {
	Dog := Dog{"Luke", 6}
	tmp := template.New("Dog Template")
	tmp, _ = tmp.Parse("Dog's name is {{.Name}} and Age is {{.Age}}")

	err := tmp.Execute(os.Stdout, Dog)
	if err != nil {
		panic(err)
	}
}