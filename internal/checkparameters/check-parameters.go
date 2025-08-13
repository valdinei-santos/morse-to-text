package checkparameters

import (
	"errors"
	"fmt"
	"path/filepath"
)

// CheckExistParameters verifica se existe ao menos 1 parametro na chamada do programa. Se não existir o programa é encerrado
func CheckExistParameters(args []string) (string, error) {
	// todosArgs := os.Args[1:]
	//if len(os.Args) == 1 {
	if len(args) == 1 {
		fmt.Println("")
		fmt.Println("Você precisa passar como parametro uma mensagem em código MORSE para ser traduzida.")
		fmt.Println("Exemplo: ./" + filepath.Base(args[0]) + " \"--- --\" ")
		fmt.Println("")
		return "", errors.New("falta o parametro na chamada do programa")
	}
	return args[1], nil
}
