// uma vez que eu recebo parâmetros variáveis, dentro da função eu trato eles com o ARRAY

package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados { // array
		fmt.Printf("%d) %s\n", i+1, aprovado) // Não sei se você já anotou isso, mas o Printf, serve para formatações, ou seja, toda vez que precisar printar algum valor que se pega da algum lugar, usar ele
	}
}
func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"} // é um slice pq ele não definiu o tamanho
	imprimirAprovados(aprovados...)
}

/* resumindo da forma que eu entendi é:

quando você usar uma função variática, na hora de usar o slice, ele deve ser colocado (reticências ...) junto na hora que você estiver digitando a função, dentro do parâmetro

igual o exemplo mostrado */
