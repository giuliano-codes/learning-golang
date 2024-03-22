package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Fornecedor struct {
	Nome    string
	Address Endereco
}

func main() {

	giuliano := Cliente{
		Nome:  "Giuliano",
		Idade: 27,
		Ativo: true,
	}

	empresa := Fornecedor{
		Nome: "giuliano.codes",
	}

	giuliano.Cidade = "Porto Alegre"
	fmt.Printf("%s\n", giuliano.Cidade)

	giuliano.Endereco.Cidade = "Bom Princípio"
	fmt.Printf("%s\n", giuliano.Cidade)

	empresa.Address.Cidade = "Porto Alegre"
	fmt.Printf("%s\n", empresa.Address.Cidade)
}
