//  DESAFIO
//
//criar um array, criar 2 slices e tentar provar que o array interno se trata do mesmo

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var itens [10]float64
	fmt.Println(itens)

	itens[0], itens[1], itens[2] = 1, 2, 3 // alterando o array, uma PARTE só
	fmt.Println(itens)

	a1 := itens[0:3]
	fmt.Println(a1) // printando só uma PARTE do array

	fmt.Println(itens)

	fmt.Println(reflect.TypeOf(a1)) // [] = slice, boa mlk, eu acho que você conseguiu criar um slice(pedaço) de um array e aqui você confirmou

}
