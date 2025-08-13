#!/bin/bash
# Script feita para gerar executavel da aplicação com o nome definido na variável

NOME_EXECUTAVEL=morse-to-text
NOME_FILE_MAIN=cmd/morse-to-text/main.go

go build -o $NOME_EXECUTAVEL $NOME_FILE_MAIN

exit 0