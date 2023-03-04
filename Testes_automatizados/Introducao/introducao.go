package main

import (
	"fmt"
	"strings"
)

func main (){
	tipoEndereco := TipoEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
	
}

// TipoEndereco verifica se o endereço tem um tipo válido e o retorna
func TipoEndereco (endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavraDoEndereco{
			return strings.Title(primeiraPalavraDoEndereco)
		}
	}
	return "Tipo inválido"
}