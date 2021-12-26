package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(numeros)

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(totalDaSoma)

	totalDaSoma2 := soma()
	fmt.Println(totalDaSoma2)

	escrever("Ola Mundo", 29, 34, 2, 7, 8, 4, 5)
}
