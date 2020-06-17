package main

import "fmt"

// nao existe operador ternario em go

func obterResultadoNota(nota float64) string {
	// return nota >=6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultadoNota(5.8))
}
