package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2           // enviando
	fmt.Println(<-ch) // recebendo
}

// dentro do chanel pode trafegar dados dentro dele

/* Podemos pensar em channels como sendo uma espécie de tunel de comunidação entre goroutines
, onde uma goroutine consegue enviar informações para outra antes mesmo de terminar sua execução.
Nesse mesmo cenário, a goroutine que recebe a informação, ficaria pausada até as informações chegarem */

// o canal é um mecanismo de SINCRONISMO

// um exemplo ele só vai continuar se um inteiro chegar no canal, até então ele não vai prosseguir

// é como se fosse uma porteira, só vai prosseguir quando aquele tipo de dado que você especficou chegar
