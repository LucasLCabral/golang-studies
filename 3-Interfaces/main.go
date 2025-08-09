package main

import "fmt"

type Pessoa interface {
	Ativar()
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("Usuario desativado.")
}

func (c *Cliente) Ativar() {
	c.Ativo = true
	fmt.Println("Usuario ativado.")
}

func Ativacao(pessoa Pessoa) {
	pessoa.Ativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	Lucas := Cliente{
		Nome:  "Lucas",
		Idade: 19,
		Ativo: false,
	}

	Ativacao(&Lucas)
	Desativacao(&Lucas)
	
}
