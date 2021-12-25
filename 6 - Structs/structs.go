package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Nome"
	u.idade = 30
	fmt.Println(u)

	endereco := endereco{"Rua 1", 100}

	u2 := usuario{"Nom 2", 32, endereco}
	fmt.Println(u2)

	u3 := usuario{idade: 34}
	fmt.Println(u3)
}
