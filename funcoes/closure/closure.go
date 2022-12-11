// a tradução de closure, seria FECHAMENTO, algo que envolve alguma coisa
// normalmente usados, quando uma função envolve outra função, algo que envolve a outra
// e isso faz com que a função INTERNA(a que seja envolvida) ela tem memória, ela tem ciência do local onde ela foi declarada
// a função EXTERNA, funciona como closure, como fechamento, ela agrupa o que ta dentro

package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}
func main() {
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()
}

// percebe-se que tem 2 variáveis chamadas de x, porém as duas vão ser executadas, normalmente
// pq isso? porque o closure, aloca memória, e dentro da função interna, existe uma memória expecífica, e o x do closure, não interfere no x da função main
// ele sabe daonde veio aquele x específico dentro daquele escopo, por isso não vai inteferir
