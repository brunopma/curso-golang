package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (itereção %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Bruno", "Fala tu", 3)
	// fale("Joao", "Só posso falar depois de vc", 1)

	// go fale("Maria", "Ei....", 500)
	// go fale("Joao", "Opa....", 500)

	go fale("Maria", "Entendi!!", 5)
	fale("Joao", "Parabens!!", 10)
}
chan type