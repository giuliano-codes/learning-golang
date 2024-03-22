package main

import "fmt"

type ID int

func main() {
	var id ID = 1

	fmt.Printf("Value of variable id: %v\n", id)
	fmt.Printf("Type of variable id: %T\n", id) // Type of variable g: float64
}
