// Vamos ver a possibilidade de definir nomes no retornos e a posição deles

package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) { // ele sempre vai retornar na ORDEM que você definiu aqui
	segundo = p2
	primeiro = p1
	return // retorno limpo, aqui só ta o return padrão porque os valores já estão setados, não preciso escrever
	// return segundo, primero

}
func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
