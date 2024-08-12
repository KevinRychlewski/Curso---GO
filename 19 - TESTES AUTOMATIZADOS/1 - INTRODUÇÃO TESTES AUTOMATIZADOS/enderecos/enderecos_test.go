// TODO ARQUIVO DE TEST EM GO TEM QUE SER O NOME DO ARQUIVO_TEST
// TESTE DE UNIDADE

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	// TestXxxxxXxxXxxx

	cenariosDeTeste := []cenarioDeTeste {
		{ "rua ABC", "rua"},
		{ "avenida Paulista", "avenida"},
		{ "rodovia dos Imigrantes", "rodovia"},
		{ "Praça das Rosas", "Tipo inválido"},
		{ "estrada da Figueira", "estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
			 TipoDeEnderecoRecebido,
			 cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}