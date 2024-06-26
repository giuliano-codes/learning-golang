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
	tamanho, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	f.Close()

	//reading files

	arquivo, err := os.ReadFile("arquivo.txt") //le bytes
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// reading chunks

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
