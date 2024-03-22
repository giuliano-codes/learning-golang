package main

import (
	"criando-modulo/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	// a := matematica.soma(10, 20) letra minuscula é privado.
	fmt.Println("Resultado: ", s)
}
