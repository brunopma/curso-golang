package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() { // metodo com receiver (receptor)
	f.turboLigado = true
}

func (f *ferrari) acelerar() {
	f.velocidadeAtual += 10
}

func main() {
	carro1 := ferrari{
		modelo:          "F40",
		turboLigado:     false,
		velocidadeAtual: 0,
	}
	carro1.ligarTurbo()
	/*
		metodo que nao está declarado no contrato da interface,
		mas pode ser usado pq a declaracao da variavel carro1
		foi feita direto no tipo ferrari, nao pela interface
		esportivo, como no carro 2 abaixo.
	*/
	carro1.acelerar()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	/*
		usar carro2.acelerar() nao faz sentido,
		pois o metodo acelerar nao esta definido
		no contrato da interface esportivo,
		de onde carro2 foi criado
	*/
	// carro2.acelerar()

	fmt.Println(carro1, carro2)
}
