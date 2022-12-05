package main

import "fmt"

func main() {
	// nessas poucas linhas, você vai ter que provar que dois SLICES APONTAM para um mesmo ARRAY interno
	// ele foi bem mais direto e simples

	s1 := make([]int, 10, 20) // lembrando que quando você criar um SLICE, dessa forma, ele automaticamente já cria um ARRAY INTERNO
	s2 := append(s1, 1, 2, 3) // percebe-se quando, ele adiciona, mais 3 números ele coloca o S1, antes das sentenças
	fmt.Println(s1, s2)
	s1[0] = 7 // quando ele muda o primeiro caracter desses slices, assim vai provar que se trata de um mesmo ARRAY
	fmt.Println(s1, s2)
}
