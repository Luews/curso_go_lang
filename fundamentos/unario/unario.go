package main

import "fmt"

func main() {
	x := 1
	y := 2
	// apenas postfix, depois da variável

	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)

	// fmt.Println(x == y--), porém não podemos fazer isso dentro de uma comparação ou até mesmo dentro de uma execução
	// por isso está comentada
}
