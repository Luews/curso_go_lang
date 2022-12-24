package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante, RECEBER OS DADOS, COMEÇO, ou seja, se eu não ler vai dar merda
	fmt.Println("Só depois que for lido")

}

func main() {

	c := make(chan int) // canal sem buffer

	go rotina(c) // operação bloqueante, MEIO, criando uma gourotine, thread

	fmt.Println(<-c) // operação bloqueante, FIM, recebendo os dados
	fmt.Println("Foi lido!")
	fmt.Println(<-c)   // aqui vai acontecer um 'deadlock' o nome é bastante intuitivo, mas ele vai morrer e as próximas linhas não vão ser executadas, pois ele está ESPERANDO um canal que ele nunca vai receber
	fmt.Println("Fim") // isso nunca vai ser printado pois está DEPOIS do deadlock, e o código para de ser executado a partir de tal ponto

}

//fatal error: all goroutines are asleep - deadlock!
