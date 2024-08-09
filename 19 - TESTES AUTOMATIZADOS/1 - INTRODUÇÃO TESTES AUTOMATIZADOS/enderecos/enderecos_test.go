// TODO ARQUIVO DE TEST EM GO TEM QUE SER O NOME DO ARQUIVO_TEST
// TESTE DE UNIDADE

package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// TestXxxxxXxxXxxx
	enderecoParaTeste := "avenida paulista"
	tipoDeEnderecoEsperado := "avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", 
		tipoDeEnderecoEsperado,
		tipoDeEnderecoRecebido,)
	}
	
}