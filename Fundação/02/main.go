package main

import "fmt"

const a bool = true

var b bool
var c int
var d string
var e float64

func main() {
	f := 3.0 // inferência
	var g float64 = 3.0

	println(a) // true
	println(b) // false
	println(c) // 0
	println(d) //
	println(e) // +0.000000e+000
	println(f) // +3.000000e+000
	println(g) // +3.000000e+000

	fmt.Printf("Type of variable g: %T\n", f) // Type of variable g: float64
}
