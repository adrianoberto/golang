package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Joao", "Marcos", "Rogerio"}

	for val := range nomes {
		fmt.Println(val)
	}

	for _, val := range nomes {
		fmt.Println(val)
	}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Rogerio",
		"sobrenome": "Oliveira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// looping infinito
	for {
		fmt.Println("Executando")
		time.Sleep(time.Second)
	}
}