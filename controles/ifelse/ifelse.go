package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota) // você sempre vai ter que colocar os Ifs e Elses, separando eles por blocos em chaves
	} // em chaves, sempre sempre sempre, separar em blocos de chaves os IFs e Elses
}
func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
