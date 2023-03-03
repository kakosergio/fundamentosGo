package enderecos

import "strings"

// TipoDeEndereco verifica se o endereço tem um tipo válido e o retorna
func TipoDeEndereco (endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavraDoEndereco{
			return primeiraPalavraDoEndereco
		}
	}
	return "Tipo inválido"
}