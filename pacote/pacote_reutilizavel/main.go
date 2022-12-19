package main

import (
	"fmt"

	"pkg/mod/github.com/cod3rcursos/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	// fmt.Println(area._TrianguloEq(5.0, 2.0)) || aqui você não vai conseguir acessar porque você viu que era privado aquela função dentro daquele pacote
}

// era pra pegar de outro pacote, mas você entendeu o conceito, só o diretório passado que está desatualizado
// você cria outro pacote em github.com, e de lá cria um repositório que exista no github mesmo, depois você, só consegue usar as funções e variáveis daquele código que você ta puxando
// no import no seu código e colocar o nome da pasta que você quer "." a função que você quer
