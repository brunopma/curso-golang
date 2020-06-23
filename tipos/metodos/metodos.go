package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string { // metodo com receiver
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) { // passando com ponteiro pra mudar diretamente a struct. nao somente seu valor
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{
		nome:      "Bruno",
		sobrenome: "Perez",
	}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Lara Viana")
	fmt.Println(p1.getNomeCompleto())
}
