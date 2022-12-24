/* Em ciência da computação, buffer de dados é uma região de memória física
utilizada para armazenar temporariamente os dados enquanto eles estão sendo movidos de um lugar para outro */

// de maneiro resumida, é uma limitação de canais, imagine-se que exista um chanel de 10 buffers, quando as 10 posições ficarem cheias, os próximos dados que vierem vão ficar em ESPERA/NA FILA

package main

import "fmt"

func rotina(ch chan int) { // Percebe-se que depois da linha 15 o buffer vai estar cheio, e as próximas linhas de código dentro desse canal, não vão ser executadas, A MENOS, que na sáida alguem consuma aquele dado, pro próximo dado entrar

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!!")
	ch <- 6

}

func main() {

	ch := make(chan int, 3) // o número 3 representa o número de buffers que esse canal vai ter de teto

	go rotina(ch)

	fmt.Println(<-ch)
}

// O QUE GERA DEADLOCK É VOCÊ FAZER UM CANAL ESPERAR POR UM DADO QUE NUNCA VAI CHEGAR
// VOCÊ ENVIAR VÁRIOS DADOS, COLOCAR UM BUFFER, E NÃO UTILIZA-LOS, NÃO CAUSA
// FALTANDO = DEADLOCK, SOBRANDO = TA SUAVE
