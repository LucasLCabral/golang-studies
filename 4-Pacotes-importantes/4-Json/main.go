package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // tageando o struct
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	// res = {"Numero":1,"Saldo":100}
	var contaX Conta
	err = json.Unmarshal(res, &contaX)
	if err != nil {
		println(err)
	}
	fmt.Printf("Conta: %d | Saldo: %d\n", contaX.Numero, contaX.Saldo)


	
	jsonPuro := []byte(`{"n":2, "s":400}`) // passando por tag

	var contaY Conta
	err = json.Unmarshal(jsonPuro, &contaY)
	if err != nil {
		println(err)
	}
	fmt.Printf("Conta: %d | Saldo: %d \n", contaY.Numero, contaY.Saldo)
}