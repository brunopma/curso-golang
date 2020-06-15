package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador a partir da atribuicao da variavel

	area := PI * m.Pow(raio, 2) //declarando e atribuindo
	fmt.Println("A área da circunferencia é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "string" // forma reduzida de criar e atribuir uma variavel (observar que é criado como var)

	fmt.Println(g, h, i)

	g = 3
	fmt.Print(g)
}
