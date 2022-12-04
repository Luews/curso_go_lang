package main

import "fmt"

func main() {
	var a int
	var b float64 // nesse caso eu preciso dizer o tipo de float, já que eu não to atribuindo um valor inicial
	var c bool
	var d string
	var e *int // *int, é um ponteiro de um inteiro

	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	// 0 0 false  <nil>, o a = 0, b = 0, c = false, d = "(nada, vazio)", e = vai ser nulo, porem no go é nil
	// esses são os valores 0 da variáveis básicas, os valores que começam

}
