// No GO, você não pode criar uma variável no código e não usar, ela é OBRIGATÓRIO, você usar a variável que criou
// se você colocar um underline, antes do import você anula ele, e se você colocar um caracter antes do import você renomeia ele

package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador, automaticamente, mas é bom saber que tipo de variável você está lidando

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2) // sempre colocar a uma variavel nova ":=" pra você declarar e já iniciarlizar ela
	fmt.Println("A área da circunferência é", area)
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)
	var e, f bool = true, false
	fmt.Println(e, f)
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
