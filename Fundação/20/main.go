package main

type Number interface {
	~int | ~float64
}

type MyNumber int

func Soma(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaGeneric[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaGenericConstraint[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func Compara2[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Giuliano": 5000, "João": 1000, "Maria": 1000}
	m2 := map[string]float64{"Giuliano": 5000, "João": 1000, "Maria": 1000}
	m3 := map[string]MyNumber{"Giuliano": 5000, "João": 1000, "Maria": 1000}

	println(Soma(m))

	println(SomaGeneric(m))
	println(SomaGeneric(m2))

	println(SomaGenericConstraint(m))
	println(SomaGenericConstraint(m2))

	println(SomaGenericConstraint(m3))

	println(Compara(5, 5.0))
	println(Compara2(5, 5.0))

}
