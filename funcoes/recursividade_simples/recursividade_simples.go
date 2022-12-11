package main

import "fmt"

func fatorial(n uint) uint { //uint = numeros inteiros positivos e sem sinais
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}
func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}

// dessa forma simples, vamos tratar apenas o resultado positivo, e te houver algum negativo, ele simplesmente não vai responder
