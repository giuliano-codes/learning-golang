package main

import "fmt"

func main() {
	var minhaVar interface{} = "Giuliano"

	println(minhaVar) // não sabe o tipo (0x730980,0x752d28)
	println(minhaVar.(string))

	resultado, okay := minhaVar.(int)

	//resultado2 := minhaVar.(int) // irá causar panic error.
	// println(resultado2) // panic: interface conversion: interface {} is string, not int

	fmt.Printf("variável resultado = %v e o resultado de okay é %v", resultado, okay)
}
