package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]
	
	// Split: separa o slice em partes, cada vez que ele encontra um espaço
	// RUA DOS BOBOS
	// ["RUA", "DOS", "BOBOS"]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo inválido"
}