package main

import "testing"

type cenarioTeste struct {
	enderecoInserido  string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenarioTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada Qaulquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenarioTeste{
		tipoDeEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		tipoRetornado := cenario.retornoEsperado
		if tipoDeEnderecoRecebido != tipoRetornado{
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, tipoRetornado)
		}
	}

}
