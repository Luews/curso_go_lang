package main

import "fmt"

// Não tem operador ternário

// return nota >= 6 ? "Aprovado" : "Reprovado" , assim seria se existisse operador ternario em Go

func obterResultado(nota float64) string { // nota que vai ser um float64, numero quebrado, e retornar em string

	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado" // aqui você retorna fora do bloco, se colocar um else, dentro do mesmo bloco ele vai reclamar, então coloque fora porque se não rodar lá dentro, vai rodar aqui fora o reprovado
}
func main() {
	fmt.Println(obterResultado(6.8))
}
