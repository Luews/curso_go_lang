// O Println, serve para printar e quebrar a linha, sendo assim o próximo comando vai ser na linha debaixo
// O %v é um parametro geralzao, mas funciona na maioria dos casos

package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x) // aqui criando uma variável em forma de string, pra não dar erro

	fmt.Println("O valor de x é " + xs)   // printando de uma forma de string
	fmt.Println("O valor de x é", x)      // aqui como se foi criada a variável xs, ele já associa e vai direto, como se fosse string
	fmt.Printf("O valor de x é %.2f.", x) // aqui é a forma correta e provavelmente você vai usar para printar floats, que você já conhecia do python, PrintF é um print formatador

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}

/*

%d = inteiro
%f = float
%.1f = float com apenas uma casa depois da vírgula
%t = valor booleano
%s = string

*/
