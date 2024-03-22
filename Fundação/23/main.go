package main

func main() {
	a := 10
	b := 2
	c := 30

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a || b < a {
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println(1)
	case 2:
		println(2)
	default:
		println("Default")
	}
}
