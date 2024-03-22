package main

import "fmt"

func main() {
	salarios := map[string]int{"Giuliano": 5000, "Carlos": 1500}

	println(salarios["Giuliano"])
	delete(salarios, "Carlos")
	salarios["Carlinhos"] = 5000
	println(salarios["Giuliano"])

	dados := make(map[string]int)
	outros_dados := map[string]int{}

	println(dados)
	println(outros_dados)

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é de %d\n", nome, salario)
	}

	for _, salario := range salarios { // _ blank identifier
		fmt.Printf("O salário é de %d\n", salario)
	}

}
