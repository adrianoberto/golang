package main

import (
	"errors"
	"fmt"
)

func main() {

	// ** inicio inteiros **
	// int - (quando criado como int, usa-se a arquitetura do computador 32 ou 64)
	var numero int32 = 100000000
	fmt.Println(numero)

	// não pode ser negativo
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias (ex: Int32 = rune)
	var numero3 rune = 123456
	fmt.Println(numero3)

	//  byte = int8
	var numero4 byte = 123
	fmt.Println(numero4)

	// ** fim inteiros **

	// ** inicio reais **
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1287878973.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123878798798.45
	fmt.Println(numeroReal3)

	// ** fim reais **

	// ** inicio string **
	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	// ** fim string **

	// ** inico booleano **

	var booleano1 bool = false
	fmt.Println(booleano1)

	// ** fim booleano **

	// ** inicio error **

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

	// ** fim error **
}
