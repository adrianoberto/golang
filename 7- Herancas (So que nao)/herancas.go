package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Joao", "Paulo", 20, 183}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciencias", "PUC"}
	fmt.Println(e1)

	fmt.Println("idade:", e1.idade)
}
