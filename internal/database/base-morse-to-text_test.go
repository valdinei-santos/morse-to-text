package database

import (
	"testing"
)

func TestBuildBaseMorseToText(t *testing.T) {
	// Caso sucesso e erro: Verifica se o mapa é construído corretamente.

	BuildBaseMorseToText()

	if baseMorseToText == nil {
		t.Fatalf("O mapa baseMorseToText não pode ser nil após a chamada de BuildBaseMorseToText.")
	}

	expectedSize := 36
	if len(baseMorseToText) != expectedSize {
		t.Errorf("Esperava que o mapa tivesse o tamanho %d, mas tem %d.", expectedSize, len(baseMorseToText))
	}

	testCases := []struct {
		chave string
		valor string
	}{
		{".-", "A"},
		{"--.", "G"},
		{"---", "O"},
		{"-----", "0"},
		{"----.", "9"},
	}

	for _, tc := range testCases {
		if val, ok := baseMorseToText[tc.chave]; !ok || val != tc.valor {
			t.Errorf("Chave '%s': esperava valor '%s', mas encontrou '%s'.", tc.chave, tc.valor, val)
		}
	}
}

func TestGetLetra(t *testing.T) {
	// Caso de sucesso e erro: Verifica o comportamento da função GetLetra.

	// Construir o mapa
	BuildBaseMorseToText()

	// Cenário 1: Chave existente (sucesso)
	t.Run("Chave Existente", func(t *testing.T) {
		chave := "..."
		expectedLetra := "S"
		letra, err := GetLetra(chave)

		if err != nil {
			t.Errorf("Não esperava erro, mas recebeu: %v", err)
		}

		if letra != expectedLetra {
			t.Errorf("Esperava a letra '%s' para a chave '%s', mas recebeu '%s'.", expectedLetra, chave, letra)
		}
	})

	// Cenário 2: Chave inexistente (erro)
	t.Run("Chave Inexistente", func(t *testing.T) {
		chave := "---..." // Uma chave que não existe no mapa
		letra, err := GetLetra(chave)

		if err == nil {
			t.Errorf("Esperava um erro para a chave '%s', mas recebeu nil.", chave)
		}

		if letra != "" {
			t.Errorf("Esperava uma string vazia, mas recebeu '%s'.", letra)
		}

		expectedErrorMsg := "chave não existe no MAP"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Esperava mensagem de erro '%s', mas recebeu '%s'.", expectedErrorMsg, err.Error())
		}
	})
}
