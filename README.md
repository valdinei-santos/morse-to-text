# Morse to Text

## Descrição Técnica

O projeto **Morse to Text** é uma aplicação CLI escrita em Go que traduz códigos Morse para texto. 
Os códigos Morse devem seguir o padrão de 1 espaço em branco para separar os códigos que representam as letras e 3 espaços em branco para separar os códigos que representam palavras.
```
Por Exemplo:
".... . -.--   .--- ..- -.. .." traduzido é HEY JUDI
```
#### Alfabeto e os respectivos códigos Morse

![Imagem Codigo Morse](images/Codigos-Morse.png)


### Estrutura do Projeto

- **cmd/morse-to-text**: Contém o ponto de entrada principal do programa (`main.go`).
- **internal/checkarguments**: Gerencia a validação e o processamento dos argumentos de entrada.
- **internal/database**: Contém a base de dados de tradução de código Morse para texto.
- **internal/translator**: Implementa a lógica de tradução de código Morse para texto.
- **images**: Contém as imagens usadas na documentação do projeto.

### Fluxo de Execução

1. O programa recebe um código Morse como argumento via linha de comando.
2. O pacote `checkarguments` valida os argumentos fornecidos.
3. O pacote `database` inicializa a base de dados de tradução.
4. O pacote `translator` realiza a tradução do código Morse para texto.
5. O resultado é exibido no terminal.

---

## Instruções para Rodar o Programa

### Pré-requisitos

- Sistema operacional Linux 
- Go instalado (versão 1.25 ou superior).

### Passos

1. Clone o repositório:
   ```sh
   git clone https://github.com/valdinei-santos/morse-to-text.git
   cd morse-to-text
   ```

2. Compile o programa:
   ```sh
   ./gera-executavel.sh
   ```

3. Execute o programa passando um código Morse como argumento:
   ```sh
   ./morse-to-text "--- .."

   Exemplo de saída:
     Código morse: --- ..
     Tradução....: OI
   ```

---

## Instruções para Rodar os Testes

### Passos

1. Navegue até o diretório do projeto:
   ```sh
   cd morse-to-text
   ```

2. Execute todos os testes:
   ```sh
   go test ./...
   ```

   Exemplo de saída:
   ```sh
   ok      github.com/valdinei-santos/morse-to-text/cmd/morse-to-text      0.024s
   ok      github.com/valdinei-santos/morse-to-text/internal/checkarguments        0.028s
   ok      github.com/valdinei-santos/morse-to-text/internal/database      0.023s
   ok      github.com/valdinei-santos/morse-to-text/internal/translator    0.023s
   ```

3. Para rodar um teste específico, use o comando:
   ```sh
   go test -run TestNomeDoTeste ./caminho/do/pacote
   ```

### Estrutura de Testes
O projeto inclui testes automatizados para os seguintes componentes:

- **checkarguments**: Valida os argumentos de entrada.
- **database**: Testa a construção da base de dados e a recuperação de letras.
- **translator**: Testa a tradução de códigos Morse, incluindo casos de sucesso e erro.
- **main**: Testa o fluxo completo do programa, incluindo cenários de sucesso e falha.

Os testes garantem que o programa funcione corretamente em diferentes cenários e que erros sejam tratados adequadamente.

## Autor

Este projeto foi desenvolvido por:

*   [Valdinei Valmir dos Santos](https://github.com/valdinei-santos)
