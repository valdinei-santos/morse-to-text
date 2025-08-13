package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain_Success(t *testing.T) {
	// Caso de sucesso: Testando a tradução de um código morse válido.

	// Simulando a chamada do programa com um argumento.
	os.Args = []string{"./morse-to-text", "-... .-. .- ... .. .-.."}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	// Verificando se a saída contém a tradução esperada.
	expected := "Tradução....: BRASIL"
	if !strings.Contains(string(out), expected) {
		t.Errorf("Esperava '%s' na saída, mas recebeu: '%s'", expected, out)
	}

}

func TestMain_NoArguments(t *testing.T) {
	// Caso de erro 1: Testando a falta de argumentos

	os.Args = []string{"./morse-to-text"}

	// Foi substituido a função os.Exit no Main para possibilitar o mock.
	var capturedExitCode int
	originalExit := exit
	exit = func(code int) {
		capturedExitCode = code
		panic("exit")
	}
	defer func() {
		exit = originalExit
	}()

	defer func() {
		if r := recover(); r != nil && r != "exit" {
			// se o panic não for do exit gera outro panic
			panic(r)
		}
	}()

	oldStdout := os.Stdout
	oldStderr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	os.Stdout = wOut
	os.Stderr = wErr

	main()

	wOut.Close()
	wErr.Close()

	out, _ := io.ReadAll(rOut)
	errOut, _ := io.ReadAll(rErr)

	os.Stdout = oldStdout
	os.Stderr = oldStderr

	if capturedExitCode != 1 {
		t.Errorf("Esperava código de saída 1, mas recebeu %d", capturedExitCode)
	}

	expectedErrorMsg := "Você precisa passar como argumento um código MORSE para ser traduzido."
	if !strings.Contains(string(errOut), expectedErrorMsg) {
		t.Errorf("Esperava '%s' na saída de erro, mas recebeu: '%s'", expectedErrorMsg, errOut)
	}

	if len(out) > 0 {
		t.Errorf("Saída padrão deveria estar vazia, mas recebeu: '%s'", out)
	}
}
