// Armazenar funções em variáveis

package main

import "fmt"

var soma = func(a, b int) int { // aqui vemos uma criação de uma variável, logo em seguida vemos uma atribuição direcionando ela a uma função, em GO podemos fazer isso
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))
	sub := func(a, b int) int { // podemos fazer a mesma coisa dentro da funcao main
		return a - b
	}
	fmt.Println(sub(2, 3))
}
