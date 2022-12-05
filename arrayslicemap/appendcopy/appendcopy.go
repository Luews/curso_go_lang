package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int // percebe-se que criou uma variável slice1 com nada nem posição
	// array1 = append(array1, 4, 5, 6), infelizmente não podemos adicionar, porque é um ARRAY != APPEND

	slice1 = append(slice1, 4, 5, 6) // aqui foram adicionados o 4 5 6, nas posições 0 1 2
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // aqui foi criado outro slice, com espaço de dois, 0 1
	copy(slice2, slice1)     // usando o método copy podemos COPIAR E COLAR os DADOS, de um slice pro outro, apenas COPIAR, sem acrescentar
	fmt.Println(slice2)      // como o tamanho desse slice é 2, 0 1, foi copiado apenas o 4 5
}
