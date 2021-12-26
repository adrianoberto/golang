package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Joel",
		"sobrenome": "da Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Antonio",
			"ultimo":   "De Paula",
		},
		"curso": {
			"nome":      "Ciencias",
			"professor": "Gilberto",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["endereco"] = map[string]string{
		"rua": "Rua 1",
	}
	fmt.Println(usuario2)

}
