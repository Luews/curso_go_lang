// uma vez que você define o tipo do array, todos os elementos vão ser daquele tipo
//tamanho fixo você define na inicialização do seu array, e ele não vai mudar
// se você colocou que tem 10 posições, ele vai permanecer com 10 posições até o final daquele array

package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10, nota-se que aqui está comentado, porque vai indo de 0...1...2... etc, e a posição que seria a 4º não existe
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ { //len para contar
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media) // para calcular a média de todos os elementos
}
