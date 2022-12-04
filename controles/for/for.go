// tem a mesma funcionalidade do While, e a mesma funcionalidade do for, porém os dois fazem o memso trabalho

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // não tem while em Go // contador, enquanto essa expressão for verdadeira o código vai continuar rodando, bem semelhante ao While
		fmt.Println(i) // euquanto i(1) for menor que dez printar o I e adicionar mais um
		i++
	}
	for i := 0; i <= 20; i += 2 { // aqui o for padrãozão, colocando variáveis e inilizando
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}
	for {
		// laço infinito // só do fato de você colocar um for sem variável nenhuma e fechar a chave o que tiver dentro da chave vai se repetir pra sempre
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
	// Veremos o foreach no capítulo de array
}
