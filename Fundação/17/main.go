package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Giuliano Arnhold"
	fmt.Printf("O cliente %v andou.", c.nome)
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) simularPonteiro(valor int) int {
	c.saldo = 5000 // (*c).saldo = 5000

	return c.saldo
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	giuliano := Cliente{
		nome: "Giuliano",
	}
	giuliano.andou()

	fmt.Printf("\nValor do nome na struct é %v.\n", giuliano.nome)

	conta := Conta{saldo: 100}
	println(conta.simular(2000))
	println(conta.saldo)

	println(conta.simularPonteiro(200))
	println(conta.saldo)

	fmt.Printf("\nO tipo da nova conta é %T.\n", NewConta())
}
