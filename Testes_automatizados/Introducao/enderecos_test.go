package main

import "testing"

func TestTipoDeEndereco(t *testing.T){
	endereco := "Rua Paulista"

	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoEndereco(endereco)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado{
		t.Errorf("O tipo %s é diferente do tipo %s!", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}

}