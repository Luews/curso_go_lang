// o defer é uma função, ela atrasa, você bota essa palavra reservada antes da sentença de código, quando você quer ATRASAR, essa execução]

package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}
func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}

//resumindo: ele atrasa o código pra você, e executa no final do escopo do código, quando a chave fecha, entre os returns
