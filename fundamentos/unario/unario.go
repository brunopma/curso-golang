package main

import "fmt"

func main() {
	x := 1
	y := 2

	// go tem apenas forma posfixada de operar em variaveis
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y + 1
	fmt.Println(y)
}
