package checkarguments

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGetArguments_WithValidArgument(t *testing.T) {
	// Caso de sucesso: Chamada da função com slice de strings com parametro.

	expectedMorseCode := "--- --"
	args := []string{"./morse-to-text", expectedMorseCode}
	morseCode, err := GetArguments(args)
	if err != nil {
		t.Errorf("Não esperava erro, mas a função retornou: %s", err)
	}

	if morseCode != expectedMorseCode {
		t.Errorf("Esperava: %s, mas a função retornou: %s", expectedMorseCode, morseCode)
	}
}

func TestGetArguments_NoArguments(t *testing.T) {
	// Caso de erro1: Chamada da função com slice de strings vazio, sem parametro.

	args := []string{"./morse-to-text"}
	morseCode, err := GetArguments(args)
	if err == nil {
		t.Errorf("Esperava um erro, mas a função retornou nil")
	}

	// Deve retornar vazio
	if morseCode != "" {
		t.Errorf("Esperava uma string vazia, mas recebeu: %s", morseCode)
	}

	expectedErrMsg := "Você precisa passar como argumento um código MORSE para ser traduzido."
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Esperava que a mensagem de erro tivesse '%s', mas recebeu: '%s'", expectedErrMsg, err.Error())
	}

	expectedExample := "./" + filepath.Base(args[0]) + " \"--- --\""
	if !strings.Contains(err.Error(), expectedExample) {
		t.Errorf("Esperava que a mensagem de erro tivesse o exemplo '%s', mas recebeu: '%s'", expectedExample, err.Error())
	}
}
