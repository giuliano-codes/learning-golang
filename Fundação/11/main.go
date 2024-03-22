package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	giuliano := Cliente{
		Nome:  "Giuliano",
		Idade: 27,
		Ativo: true,
	}

	giuliano.Nome = "Giuliano Arnhold"

	fmt.Printf("%s", giuliano.Nome)
}
