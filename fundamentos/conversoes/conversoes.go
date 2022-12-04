package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2                      // eu poderia também converter o x para int e assim fazer, mas não daria o resultado exato, para uma divisão, subtração, etc, as duas variáveis precisam estar no mesmo tipo
	fmt.Println(x / float64(y)) // eu preciso converter as duas variáveis para o mesmo tipo, se não a função não funciona, nesse caso eu fiz as duas com o mesmo tipo (float64)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) // aqui transformei a nota final em um inteiro e printei, o go não aredondou pra 7

	// cuidado...
	fmt.Println("Teste " + string(rune(97))) // você não pode simplesmente fazer colocar dessa forma, você precisa converter para uma string, se não vai printar o tipo, bizarro, e meio complexo mas fodas-se

	// int para string (strconv.Itoa), int to (a) string
	fmt.Println("Teste " + strconv.Itoa(123)) // strconv, sintaxe para converter, você vai usar muito ela

	// string para int (strconv.Atoi), ao contrário, string (a) to int
	num, _ := strconv.Atoi("123") // nesse caso ele vai te retornar 2 variáveis, e como o GO não pode ficar variável sem usar, você anula a 2º usando o underline
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // boolean para string, tudo que não for 1 ou true, vai retornar false, então tipo, não tem segredo
	if b {
		fmt.Println("Verdadeiro")
	}
}
