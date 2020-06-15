package main

import "fmt"

func main() {
	fmt.Print("texto 1")
	fmt.Print("texto 2")
	fmt.Println("texto 3")
	fmt.Println("texto 4")

	PI := 3.14159

	fmt.Println("O valor de x é:", PI)
	fmt.Println("O valor de x é: " + fmt.Sprint(PI))
	fmt.Printf("O valor de x é: %.2f.", PI)

	const (
		a = 1
		b = 1.9999
		c = false
		d = "string"
	)

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
