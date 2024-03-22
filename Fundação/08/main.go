package main

import (
	"errors"
	"fmt"
)

func main() {
	println(sum(1, 2))
	valor, err := sum_multiple_returns_error(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a int, b int) int {
	return a + b
}

func soma(a, b int) int {
	return a + b
}

func sum_multiple_returns(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

func sum_multiple_returns_error(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
