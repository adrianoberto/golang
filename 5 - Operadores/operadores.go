package main

import "fmt"

func main() {
	// Aritimeticos (Padrões)

	soma := 1 + 1
	divisao := 1 / 1
	subtracao := 1 - 1
	multiplicacao := 1 * 1
	restoDivisao := 1 % 1
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// FIm Aritimeticos

	// Atribuição
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)
	// Fim Atribuição

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//fim Operadores relacionais

	// Operadores Logicos
	fmt.Println("--------------------------------")
	fmt.Println("Operadores Logicos")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro, !falso)
	fmt.Println("--------------------------------")
	// FIM Operadores Logicos

	// Operadores unarios
	fmt.Println("Operadores unarios")
	numero6 := 10
	numero6++
	fmt.Println(numero6)
	numero6 += 15
	fmt.Println(numero6)
	numero6--
	fmt.Println(numero6)
	numero6 -= 15
	fmt.Println(numero6)
	numero6 *= 15
	fmt.Println(numero6)
	numero6 /= 15
	fmt.Println(numero6)
	numero6 %= 15
	fmt.Println(numero6)
	fmt.Println("--------------------------------")
	// FIM Operadores unarios

	// Operador ternario
	fmt.Println("Não tem operador ternario")
}
