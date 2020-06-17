package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      11341.00,
		"Gabriela Silva": 13351.00,
		"Pedro Junior":   15344.00,
		"Bruno Perez":    16897.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
