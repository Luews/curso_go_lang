package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo, criar uma VARIÁVEL E JÁ INICIALIZAR ELA
	i += 3 // i = i + 3, como é uma variável você pode modificar ela durante o código, como por exemplo somar
	i -= 3 // i = i - 3, subtrair, tirar
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2
	fmt.Println(i)

	x, y := 1, 2 // := , para criar e inicializar já
	fmt.Println(x, y)

	x, y = y, x // você pode trocar as variáveis de posição, o x passou a ser y, e o y passou a ser x
	fmt.Println(x, y)
}
