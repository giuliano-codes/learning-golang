package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) desativar() {
	c.Ativo = false
	fmt.Printf("%t\n", c.Ativo)
}

func main() {

	giuliano := Cliente{
		Nome:  "Giuliano",
		Idade: 27,
		Ativo: true,
	}

	fmt.Printf("%t\n", giuliano.Ativo)
	giuliano.desativar()
	fmt.Printf("%t\n", giuliano.Ativo)
}
