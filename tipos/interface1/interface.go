package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interface sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Bruno", "Perez"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"lapis", 3.99} // polimorfismo, a variavel coisa pode ser de qualquer tipo que esteja no contrato da interface
	fmt.Println(coisa.toString())
	imprimir(coisa)
}
