// utilizado para reciclagens de códigos, mas não é uma herança de verdade, porém nesse método você consegue reciclar alguns códigos

package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}
type ferrari struct { // percebe-se que o carro não tem o tipo dele, por isso tem aquela sensação que você ta herdando, mas na verdade só ta fazendo uma composição
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
