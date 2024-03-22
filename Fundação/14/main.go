package main

type Pessoa interface { // interface aceita apenas métodos
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func main() {
	giuliano := Cliente{
		Nome:  "Giuliano",
		Idade: 27,
		Ativo: true,
	}

	empresa := Empresa{
		Nome: "Dev",
	}

	Desativacao(giuliano)
	Desativacao(empresa)

}
