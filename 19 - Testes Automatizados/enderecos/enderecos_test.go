// TESTE DE UNIDADE

package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	RetornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Praça dos Meninos", "Tipo Inválido"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.RetornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.RetornoEsperado,
			)
		}
	}
}
