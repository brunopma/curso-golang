package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[23198723123] = "Ana"
	aprovados[231987312133] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[23198723123])

	delete(aprovados, 23198723123)

	fmt.Println(aprovados)
}
