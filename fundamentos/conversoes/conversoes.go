package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(123)) // 123 converte pro ascii relativo ao 123, que é {

	// int para string
	fmt.Println("Test " + strconv.Itoa(123))

	// String pra int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("false")
	if !b {
		fmt.Print("b é true")
	}
}
