package main

import "fmt"

func notaParaConceito(nota float64) string {
	switch {
	case nota <= 10 && nota > 8:
		return "A"
	case nota <= 8 && nota > 6:
		return "B"
	case nota <= 6 && nota > 4:
		return "C"
	case nota <= 4 && nota > 2:
		return "D"
	case nota <= 2 && nota >= 0:
		return "E"
	default:
		return "Nota Invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7.2))
	fmt.Println(notaParaConceito(5.4))
	fmt.Println(notaParaConceito(3.3))
	fmt.Println(notaParaConceito(1.1))
	fmt.Println(notaParaConceito(0.0))
	fmt.Println(notaParaConceito(1199.11))
	fmt.Println(notaParaConceito(-22.11))
}
