package checkarguments

import (
	"errors"
	"path/filepath"
)

// GetArguments verifica se existe ao menos 1 argumento na chamada do programa.
func GetArguments(args []string) (string, error) {
	if len(args) == 1 || args[1] == "" {
		return "", errors.New("Você precisa passar como argumento um código MORSE para ser traduzido. \nExemplo: ./" + filepath.Base(args[0]) + " \"--- --\"")
	}
	return args[1], nil
}
