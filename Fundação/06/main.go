package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("len = %d; cap = %d; value=%v \n", len(s), cap(s), s)

	fmt.Printf("len = %d; cap = %d; value=%v \n", len(s[:3]), cap(s[:3]), s[:3]) // inicio até 3 elementos

	fmt.Printf("len = %d; cap = %d; value=%v \n", len(s[3:]), cap(s[3:]), s[3:]) // pula 3 elementos até o final

	// adicionando novo valor
	s = append(s, 20)

	fmt.Printf("len = %d; cap = %d; value=%v \n", len(s), cap(s), s)
}
