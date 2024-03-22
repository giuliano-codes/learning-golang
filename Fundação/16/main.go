package main

func soma(a, b int) int {
	a += 50
	return a + b
}

func somaPonteiro(a, b *int) int {
	*a = 50

	return *a + *b
}

func main() {
	a := 10
	b := 20

	println(soma(a, b))
	println(a)

	println(somaPonteiro(&a, &b))
	println(a)
}
