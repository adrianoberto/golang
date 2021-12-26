package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor ou igual que 10")
	}

	// if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que menos 10")
	} else {
		fmt.Println("Número entre 0  e -10")
	}
}
