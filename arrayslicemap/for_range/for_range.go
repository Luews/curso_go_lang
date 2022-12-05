// Nessa aula, vamos aprender a colocar um range, um alcance para usarmos o for

package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!, o próprio compilador conta
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero) // aqui ele vai printar o i e o numero, que seria, o i ele sempre vai acrescentar mais 1, e o numero vai printar ele mesmo, com o range de numeros
	}
	for _, num := range numeros { // aqui nos podemos anular o indice e printar somente os numeros
		fmt.Println(num)
	}
}
