package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Escrevendo um arquivo com Go!"))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho escrito: %d bytes\n", tamanho)
	f.Close()

	/// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura de pouco a pouco ( buffer )
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 5) // tamanho do buffer
	
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println("Lendo do arquivo:")
		fmt.Println(string(buffer[:n]))
	}
	arquivo2.Close()

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
