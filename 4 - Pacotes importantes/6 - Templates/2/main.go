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
	t := template.Must(template.New("Dog Template").Parse("Dog's name is {{.Name}} and Age is {{.Age}}"))
	err := t.Execute(os.Stdout, Dog)
	if err != nil {
		panic(err)
	}
	
}