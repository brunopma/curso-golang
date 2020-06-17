package main

import "fmt"

func main() {
	fmt.Println(media(1.42, 3.45, 5.25, -13.42, 5.00, 9.25, 99.99))
	fmt.Println(media())
}

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}
