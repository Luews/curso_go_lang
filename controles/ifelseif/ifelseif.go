// Usando o else if, eu posso colocar quantos CAMINHOS eu QUISER, do que usar o if e else padrão
/* Dessa forma eu crio vários caminhos multiplos, mas normalmente a gente usa o Switch, fica mais amplo,
mas se for algo "pequeno" é legal colocar uns else if, porém se for um código com 200 caminhos, fica muito trabalhoso e dificil mexer até */

package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A" // O return significa que ele vai sair do método e RETORNAR esse valor
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}
func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
