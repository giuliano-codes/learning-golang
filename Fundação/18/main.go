package main

import "fmt"

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T.", t)
}

func main() {
	var x interface{} = 10
	var y interface{} = "hello, world!"

	showType(x)
	showType(y)
}
