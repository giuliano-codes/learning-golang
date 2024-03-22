package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"un", "dois", "três"}

	for k, v := range numeros {
		println(k, v)
	}

	for _, v := range numeros {
		println(v)
	}

	j := 0
	for j < 10 {
		println(j)
		j++
	}

	for {
		println("Hello, World!")
	}
}
