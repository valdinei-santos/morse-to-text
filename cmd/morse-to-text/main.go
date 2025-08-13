package main

import (
	"fmt"
	"os"

	"github.com/valdinei-santos/morse-to-text/internal/checkparameters"
	"github.com/valdinei-santos/morse-to-text/internal/database"
	"github.com/valdinei-santos/morse-to-text/internal/translator"
)

// var baseTextToMorse map[string]string
// var baseMorseToText map[string]string

func main() {
	codigoMorse, err := checkparameters.CheckExistParameters(os.Args)
	if err != nil {
		os.Exit(1)
	}
	database.BuildBaseMorseToText()
	texto, err := translator.TranslateMorseToText(codigoMorse)
	if err != nil {
		fmt.Printf("ERRO: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Código morse: %s\n", codigoMorse)
	fmt.Printf("Tradução....: %s\n", texto)
}

/* func checkExistParameters() {
	// todosArgs := os.Args[1:]
	if len(os.Args) == 1 {
		fmt.Println("")
		fmt.Println("Você precisa passar como parametro uma mensagem em código MORSE para ser traduzida.")
		fmt.Println("Exemplo: ./" + filepath.Base(os.Args[0]) + " \"--- --\" ")
		fmt.Println("")
		os.Exit(1)
	}
} */

/* func translateMorseToText(c string) string {
	texto := ""
	palavras := strings.Split(c, "   ")
	for _, palavra := range palavras {
		letras := strings.Split(palavra, " ")
		for _, letra := range letras {
			texto = texto + baseMorseToText[string(letra)]
		}
		texto = texto + " "
	}
	return texto
} */

/* func buildBaseMorseToText() {
	baseMorseToText = map[string]string{
		".-":    "A",
		"-...":  "B",
		"-.-.":  "C",
		"-..":   "D",
		".":     "E",
		"..-.":  "F",
		"--.":   "G",
		"....":  "H",
		"..":    "I",
		".---":  "J",
		"-.-":   "K",
		".-..":  "L",
		"--":    "M",
		"-.":    "N",
		"---":   "O",
		".--.":  "P",
		"--.-":  "Q",
		".-.":   "R",
		"...":   "S",
		"-":     "T",
		"..-":   "U",
		"...-":  "V",
		".--":   "W",
		"-..-":  "X",
		"-.--":  "Y",
		"--..":  "Z",
		"-----": "0",
		".----": "1",
		"..---": "2",
		"...--": "3",
		"....-": "4",
		".....": "5",
		"-....": "6",
		"--...": "7",
		"---..": "8",
		"----.": "9",
	}
} */

/* func buildBaseTextToMorse() {
	baseTextToMorse = map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"0": "-----",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
	}
} */
