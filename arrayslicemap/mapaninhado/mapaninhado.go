package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  15456.78,
			"Geraldo Pereira": 6706.78,
		},
		"J": {
			"João Neto":    14565.28,
			"João Gabriel": 4565.21,
		},
		"P": {
			"Pedro Junior": 4565.30,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
