# mfmt

Pequenos utilitários de terminal em Go que complementam o pacote `fmt` ("my formatting").

## O que faz

- `mfmt.Scanln() string`: lê uma linha inteira da entrada padrão e devolve o texto sem espaços
  e quebra de linha nas pontas. Diferente do `fmt.Scanln`, não para no primeiro espaço.
- `mfmt.Cleanse()`: limpa o terminal rodando o comando `clear`.

## Por que existe

São trechos que eu repetia em todo programa de terminal. É uma das minhas bibliotecas pessoais,
para reaproveitar nos meus projetos.

## Instalação

Pré-requisito: Go 1.25.4+ num módulo Go. O `Cleanse` depende do comando `clear` (Linux/macOS).

```sh
go get github.com/rillmind/mfmt@latest
```

Exemplo de uso:

```go
package main

import (
	"fmt"

	"github.com/rillmind/mfmt"
)

func main() {
	mfmt.Cleanse()
	fmt.Print("Seu nome: ")
	nome := mfmt.Scanln()
	fmt.Println("Olá, " + nome + "!")
}
```

## Como testar

Não há testes automatizados. `go vet ./...` no repositório confere o código. Para conferir o
comportamento, rode o exemplo acima:

```sh
echo '  Raul Holanda  ' | go run .   # a tela é limpa e aparece "Seu nome: Olá, Raul Holanda!"
```

O nome completo, com espaço no meio, é lido inteiro, e os espaços das pontas são removidos.
