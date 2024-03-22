package main

func main() {
	var meuArray [3]int // array tem tamanho fixo
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for indice, valor := range meuArray {
		println(indice)
		println(valor)
	}
}
