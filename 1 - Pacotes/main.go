package main

import (
	"fmt"
	"modulo/auxilliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxilliar.Escrever()

	erro := checkmail.ValidateFormat("meutestegmailcom")
	fmt.Println(erro)
}
