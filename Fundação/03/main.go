package main

import "fmt"

type ID int

func main() {
	var id ID = 1
	println(id)
	fmt.Printf("Type of variable g: %T\n", id) // Type of variable g: float64
}
