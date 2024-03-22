package main

import "fmt"

func main() {
	a := 10
	println(a)  // printa o valor
	println(&a) // printa o endereço da variável

	var ponteiro *int = &a
	println(ponteiro)
	println(*ponteiro) // printa o valor do endereço
	*ponteiro = 20
	println(a)
	b := &a
	fmt.Printf("Type of variable b: %T\n", b)
	println(b)
	println(*b)

}
