/* uma função é um bloco de código onde vai ter vários parametros, é como se fosse uma receita de bolo, vão ter vários
ingredientes, varios caminhos de fazer, pra no final te retornar um bolo, pode ter vários resultados, ou melhor... vários bolospackage funcoesgo
mas o importante é saber que resumidamente, vai ser um monte de código, ifs, prints, etc, compilados pra ter dar um resultado só */

package main

import "fmt"

func somar(a int, b int) int { // passando o tipo de dados tanto na entrada, quanto na saida
	return a + b
}
func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)

}
