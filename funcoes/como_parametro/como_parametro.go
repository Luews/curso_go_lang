/* Nessa aula vamos aprender a como passar uma FUNÇÃO como PARAMETRO para OUTRA FUNÇÃO

É importante deixar claro que GO não é em essência uma linguagem funcional, você vai perceber que a API de GO, por exemplo,
não trabalha de uma forma funcional, MAS, o fato de você poder ARMAZENAR função em variável, passar função como parâmetro,
todos esses recursos em GO permite que você use príncipios da programação funcional, dentro dos seus programas

Embora Go não seja uma linguagem funcional em sua essência */

// map, reducer, filter, procurar pra pegar de exemplos o porque de fazer uma função onde recebe outra função como parametro

package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
func main() {
	resultado := exec(multiplicacao, 3, 4) // percebe-se que aqui se usa as duas funções || a exec primeiro, pra servir de base pra multiplicacao se adequar
	fmt.Println(resultado)
}

/* uma ideia que eu pensei pra deixar o código mais limpo é, fazer essas funções dentro de funções como parâmetro

E fazer como por exemplo, nesse mesmo exemplo citado nessa aula:

resultado1 :- exec(divisao, 4, 5)
resultado2 := exec(subtracao, 5, 2)

ficar mais limpo, só pensando em cima desse exemplo, no futuro da pra imaginar situações onde eu posso encaixar, para ficar mais rápido e simples */
