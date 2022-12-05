// podemos criar um SLICE com o método MAKE e AUTOMATICAMENTE ele cria um ARRAY

package main

import "fmt"

func main() {

	s := make([]int, 10) // aqui vamos criar um slice (pedaço), e automaticamente ele vai criar um array interno dele
	s[9] = 12            // na posição 9 vai trocar, pelo número 12
	fmt.Println(s)

	s = make([]int, 10, 20) // aqui vamos especificar, o slice (pedaço, lembrando), tem 10 de espaço ocupando, e com a CAPACIDADE de 20(que seria o array)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // append = adicionar, e aqui vamos acrescentar mais 10 espaços, algarismos, numeros, enfim nesse slice, fazendo com que ele atinja a capacidade
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // aqui no caso adicionamos mais um, o slice não vai dar um erro, ele simplesmente vai dobrar automaticamente a capacidade e conseguir adicionar mais um
}
