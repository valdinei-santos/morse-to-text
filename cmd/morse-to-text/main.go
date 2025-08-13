package main

import (
	"fmt"
	"os"

	"github.com/valdinei-santos/morse-to-text/internal/checkarguments"
	"github.com/valdinei-santos/morse-to-text/internal/database"
	"github.com/valdinei-santos/morse-to-text/internal/translator"
)

// usada para viabilizar o mock nos testes
var exit = os.Exit

func main() {
	morseCode, err := checkarguments.GetArguments(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, "")
		exit(1)
	}
	database.BuildBaseMorseToText()
	texto, err := translator.TranslateMorseToText(morseCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRO: %s\n", err)
		exit(1)
	}
	fmt.Printf("Código morse: %s\n", morseCode)
	fmt.Printf("Tradução....: %s\n", texto)
}
