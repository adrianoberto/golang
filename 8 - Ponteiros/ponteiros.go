package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel3 int = 100
	var ponteiro *int = &variavel3
	fmt.Println(variavel3, ponteiro, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, ponteiro, *ponteiro)
}
