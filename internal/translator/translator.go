package translator

import (
	"errors"
	"strings"

	"github.com/valdinei-santos/morse-to-text/internal/database"
)

// TranslateMorseToText faz a tradução do código morse para texto.
func TranslateMorseToText(c string) (string, error) {

	texto := ""
	palavras := strings.Split(c, "   ")
	for _, palavra := range palavras {
		letras := strings.Split(palavra, " ")
		for _, letra := range letras {
			l, err := database.GetLetra(string(letra))
			if err != nil {
				return "", errors.New("o código morse " + letra + " não foi encontrado!")
			}
			texto = texto + l
		}
		texto = texto + " "
	}
	return texto, nil
}
