package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(2, 5)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resulado := f("Texto da função 1")
	fmt.Println(resulado)

	resultadosSoma, resultadosSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma, resultadosSubtracao)

	resultadosSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma2)

	_, resultadosSubtracao2 := calculosMatematicos(10, 15)
	fmt.Println(resultadosSubtracao2)
}
