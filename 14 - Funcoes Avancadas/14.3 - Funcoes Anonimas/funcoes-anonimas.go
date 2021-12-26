package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Olá mundo 2")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 89)
	}("Olá mundo 2")
	fmt.Println(retorno)

	funcao := func(texto string) {
		fmt.Println(texto)
	}
	funcao("Olá mundo 3")
}
