// Aqui ta dando erro, mas no cmd do windows, onde desativei o modulo, deu certo

package main

import (
	"fmt"
)

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	Distancia(2, 5)

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}

/* primeiro de tudo você vai ter que fazer:

toda vez que entrar no VS CODE, dar um go env;

e lembrar de alterar o GO111MODULE=off;

para continuar no curso;

set GO111MODULE=off; ISSO NO CMD DO WINDOWS */

/*

Executando Múltiplos Arquivos no Windows


No caso do Windows o comando abaixo não funciona:

go run *.go


Quando for necessário executar múltiplos arquivos no Windows, a forma correta é:

go run arquivo1.go arquivo2.go


Bons estudo! */
