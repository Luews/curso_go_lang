package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo
func doisTresQuatroVezes(base int, c chan int) {

	time.Sleep(time.Second)

	c <- 2 * base // enviando dados para o canal
	time.Sleep(time.Second)

	c <- 3 * base
	time.Sleep(3 * time.Second)

	c <- 4 * base
}
func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("Iniciando código...") // percebe-se que eu escrevi esse println depois da linha da goroutine, porém ele foi executado primeiro

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c) // aqui o ultimo canal, sendo executado de forma diferente

	fmt.Println("Código finalizado!")
}

// uma gourotine de certa forma, causa um descincronismo
// e um chanel de certa forma também, causa um sincronismo
