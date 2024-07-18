package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
	// Saldo  int `json:"-"` //omitir informação
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := []byte(`{"n":5,"s":500}`)
	var contaX Conta
	json.Unmarshal(jsonPuro, &contaX)

	println(contaX.Saldo)
	println(contaX.Numero)
}
