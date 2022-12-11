// Resumindo uma função, é um bloco de código já PROGRAMADO, para realizar algo
// Como se fosse a receita de bolo, tudo já está PROGRAMADO para executar o bolo,
// Agora, os PARÂMETROS, que pode servir de exemplo "tipo de ovo" "sabor do bolo" etc, nós vamos colocar

package main

import "fmt"

func f1() { // função sem parametro, o parametro da função você passa entre as () da sua função, no caso dessa, é SEM
	fmt.Println("F1")
}

func f2(p1 string, p2 string) { // aqui vemos uma função com parametro, ()
	fmt.Printf("F2: %s %s\n", p1, p2) // porém sem retorno, aqui nós vemos, que vamos atribuir
}

func f3() string { // sem parâmetro porém com retorno
	return "F3"
}

func f4(p1, p2 string) string { // percebe-se que se os parametros forem do mesmo tipo, você pode passar 1Milhão de parametros e colocar o tipo de todos eles no final
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) { // aqui vemos uma função com 2 tipos de retorno, no caso você vai ter que declarar as duas, ou colocar um _ se quiser usar ela, se for executar eles como variável
	return "Retorno 1", "Retorno 2" // retornos múltiplos
}

func main() {
	f1()
	f2("Param1", "Param2")
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
