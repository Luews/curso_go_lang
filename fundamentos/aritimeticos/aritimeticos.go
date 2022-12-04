package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise, bit a bit
	fmt.Println("E =>", a&b)   // 11 & 10 = 10 , em forma de binário
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01
	c := 3.0
	d := 2.0

	// outras operacoes usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // O valor maior dentre esses dois, no caso (3 ou 2)
	fmt.Println("Menor =>", math.Min(c, d))                   // O valor menor dentre esses dois, no caso (3 ou 2) também, mas no caso float64
	fmt.Println("Exponenciação =>", math.Pow(c, d))           // O valor em forma de potência,  3 elevado a 2, 3 * 3 = 9

}
