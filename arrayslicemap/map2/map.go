// nesse map2, vimos uma forma de inicializar o valor de uma forma literal(modo mais fácil eu achei) na própria declaração do map
// já declarou e já inicializou e já foi atribuindo os valores iniciais

package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{ //a chave vai ser uma string (um nome) e o valor vai ser um float64(o salário)
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, // sempre vai precisar da vírgula nos valores atribuídos, ou a chave encostada no último
	}

	funcsESalarios["Rafael Filho"] = 1350.0     // podemos adicionar mais, dessa forma
	delete(funcsESalarios, "Inexistente")       // se você tentar excluir algo que não existe, não tem problema, o Golang vai entender e vai seguir sem erros
	for nome, salario := range funcsESalarios { // forzinho padrão passando por todas as atribuições e mandando printar
		fmt.Println(nome, salario)
	}
}
