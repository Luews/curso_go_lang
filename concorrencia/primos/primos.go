package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool { // lógica padrão se você for lendo direitinho e seguir pela regra de números primos, você consegue entender
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func primos(n int, c chan int) { // a quantidade de números primos
	inicio := 2
	for i := 0; i < n; i++ { // laço que vai ser responsável por encontrar todos os números primos que eu quero
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	close(c) // precisamos fechar o canal, pode gerar um deadlock
}
func main() {
	c := make(chan int, 30)
	go primos(60, c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}

// podemos usar o laço for em cima de um canal

// podemos
