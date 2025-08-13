# Morse to Text

## Descrição Técnica

O projeto **Morse to Text** é uma aplicação escrita em Go que traduz códigos Morse para texto. Ele utiliza uma estrutura modular, com diferentes pacotes para gerenciar argumentos de entrada, a base de dados de tradução e a lógica de tradução. A aplicação segue boas práticas de organização de código e inclui testes automatizados para garantir a qualidade e a confiabilidade do sistema.

### Estrutura do Projeto

- **cmd/morse-to-text**: Contém o ponto de entrada principal do programa (`main.go`).
- **internal/checkarguments**: Gerencia a validação e o processamento dos argumentos de entrada.
- **internal/database**: Contém a base de dados de tradução de código Morse para texto.
- **internal/translator**: Implementa a lógica de tradução de código Morse para texto.
- **tests**: Testes automatizados para todas as partes do sistema, garantindo cobertura de casos de sucesso e erro.

### Fluxo de Execução

1. O programa recebe um código Morse como argumento via linha de comando.
2. O pacote `checkarguments` valida os argumentos fornecidos.
3. O pacote `database` inicializa a base de dados de tradução.
4. O pacote `translator` realiza a tradução do código Morse para texto.
5. O resultado é exibido no terminal.

---

## Instruções para Rodar o Programa

### Pré-requisitos

- Go instalado (versão 1.25 ou superior).

### Passos

1. Clone o repositório:
   ```sh
   git clone https://github.com/valdinei-santos/morse-to-text.git
   cd morse-to-text

2. Compile o programa:
   ```sh
   ./gera-executavel.sh

3. Execute o programa passando um código Morse como argumento:
   ```sh
   ./morse-to-text "--- .."

   Exemplo de saída:
   Código morse: --- ..
   Tradução....: OI

---

## Instruções para Rodar os Testes

### Pré-requisitos

- Go instalado (versão 1.25 ou superior).

### Passos

1. Navegue até o diretório do projeto:
   cd morse-to-text

2. Execute todos os testes:
   ```sh
   go test ./...

   Exemplo de saída:
   ```sh
   ok  	github.com/valdinei-santos/morse-to-text/internal/checkarguments	0.005s
   ok  	github.com/valdinei-santos/morse-to-text/internal/database		0.003s
   ok  	github.com/valdinei-santos/morse-to-text/internal/translator		0.004s

3. Para rodar os testes específicos, use o comando:
   ```sh
   go test -run TestNomeDoTeste ./caminho/do/pacote

### Estrutura de Testes
O projeto inclui testes automatizados para os seguintes componentes:

- **checkarguments**: Valida os argumentos de entrada.
- **database**: Testa a construção da base de dados e a recuperação de letras.
- **translator**: Testa a tradução de códigos Morse, incluindo casos de sucesso e erro.
- **main**: Testa o fluxo completo do programa, incluindo cenários de sucesso e falha.

Os testes garantem que o programa funcione corretamente em diferentes cenários e que erros sejam tratados adequadamente.
