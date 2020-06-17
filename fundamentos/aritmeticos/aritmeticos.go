package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Modulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^10 = 01

	c := 3.0
	d := 3.0

	// usando o pacote math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Exponenciacao =>", math.Pow(c, d))
}
