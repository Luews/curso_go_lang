package main

import (
	"fmt"
	"strings"
)

type pessoa struct { // struct normal/ classe normal,
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string { // uma função de pegar o nome completo
	return p.nome + " " + p.sobrenome
}
func (p *pessoa) setNomeCompleto(nomeCompleto string) { // percebe-se que devemos utilizar o ponteiro, para alocar uma memória e usar de referência, para assim conseguirmos MUDAR
	partes := strings.Split(nomeCompleto, " ") // o strings split vai separar os indices das strings, então, a string 0 vai ser a primeira, e a string 1 vai ser a segunda
	p.nome = partes[0]
	p.sobrenome = partes[1]
}
func main() {
	p1 := pessoa{"Pedro", "Silva"}    // criar uma variável p1 e já inicializar ela
	fmt.Println(p1.getNomeCompleto()) //padrãozão normal

	p1.setNomeCompleto("Maria Pereira") // usando a função onde foi colocado o ponteiro, criou uma locação de memória nova e usando apenas de referências aqueles argumentos
	fmt.Println(p1.getNomeCompleto())   // printando vai sair um resultado diferente, mudando conforme a função com ponteiro
}
