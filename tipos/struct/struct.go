package main

import (
	"fmt"
	"os"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 { // Método: função com receiver (receptor), fazendo com que a chama fique igual a OO
	return p.preco * (1 - p.desconto)
}

func precoComDesconto2(p produto) float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	/*
		criaçao de struct usando notaçao completa
		primeiro instanciando uma variavel e depois
		preenchendo com cada um dos valores da struct
	*/
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto(), precoComDesconto2(produto1))

	/*
		criaçao de struct usando notaçao compacta
		com parametros da struct nomeados
	*/
	produto2 := produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto2)

	/*
		criaçao de struct usando notaçao compacta
		com parametros da struct na ordem criada
	*/
	produto3 := produto{"lapis", 1.79, 0.05}
	fmt.Println(produto3)

	os.Setenv("GO_TEST_VAR", "TESTANDO123")
}
