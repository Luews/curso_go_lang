// funções variáticas, são funções que recebem quantidade de parâmetros variáveis

package main

import "fmt"

func media(numeros ...float64) float64 { // a sintaxe dessa função que é variática, são as reticências (...)
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros)) // aqui pra retornar o valor dividido, por todos as outras variáveis
}
func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}

// no final, percebe-se que a palavra numeros, retorna VÁRIOS numeros quebrados(float64), e isso só é possivel por conta das reticências

//a palavra numeros aloca várias outras variáveis
