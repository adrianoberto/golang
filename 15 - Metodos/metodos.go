package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (user usuario) maiorDeIdade() bool {
	return user.idade > 18
}

func (u *usuario) fazendoAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario1", 30}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Usuario 2", 45}
	usuario2.salvar()

	maiorDeIdate := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdate)

	usuario2.fazendoAniversario()
	fmt.Println(usuario2.idade)
}
