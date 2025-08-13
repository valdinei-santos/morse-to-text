package translator

import (
	"testing"

	"github.com/valdinei-santos/morse-to-text/internal/database"
)

func TestTranslateMorseToText(t *testing.T) {
	database.BuildBaseMorseToText()

	// Cenário 1: Tradução de um código morse válido (múltiplas palavras).
	t.Run("Traducao valida com mais de uma palavra", func(t *testing.T) {
		morseCode := "--- .-.. .-   -- ..- -. -.. ---"
		expectedText := "OLA MUNDO "

		translatedText, err := TranslateMorseToText(morseCode)

		if err != nil {
			t.Errorf("Não esperava um erro, mas recebeu: %s", err)
		}

		if translatedText != expectedText {
			t.Errorf("Esperava '%s', mas recebeu '%s'", expectedText, translatedText)
		}
	})

	// Cenário 2: Código morse válido (uma única palavra).
	t.Run("Traducao valida com uma palavra", func(t *testing.T) {
		morseCode := "-... .-. .- ... .. .-.."
		expectedText := "BRASIL "

		translatedText, err := TranslateMorseToText(morseCode)

		if err != nil {
			t.Errorf("Não esperava um erro, mas recebeu: %s", err)
		}

		if translatedText != expectedText {
			t.Errorf("Esperava '%s', mas recebeu '%s'", expectedText, translatedText)
		}
	})

	// Cenário 3: Código morse inválido (letra que não existe).
	t.Run("Traducao com codigo invalido", func(t *testing.T) {
		morseCodeInvalido := ".... --- .-. .- --..... --.." //  "--....." não existe.
		_, err := TranslateMorseToText(morseCodeInvalido)
		if err == nil {
			t.Errorf("Esperava um erro, mas recebeu nil")
		}

		// A mensagem de erro deve conter o código morse inválido.
		expectedErrorMsg := "o código morse --..... não foi encontrado!"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Esperava mensagem de erro: '%s', mas recebeu: '%s'", expectedErrorMsg, err.Error())
		}
	})

}
