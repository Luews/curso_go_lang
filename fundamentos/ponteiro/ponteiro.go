// O ponteiro nada mais é do que uma referência de endereço do memória
// Nós vamos criar um ponteiro da anotação, variável "*", do lado do tipo de ponteiro

// Resumindo, Lucas do futuro, o ponteiro é tipo quando você ta no youtube e você quer comentar um momento expecífico do vídeo, você usa "6:32 mds que cringe", é exatamente isso, só que você vai usar nos códigos em Go daqui pra frente

package main

import "fmt"

func main() {
	i := 1
	var p *int = nil
	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++
	// Go não tem aritmética de ponteiros
	// p++
	fmt.Println(p, *p, i, &i)
}

// as duas variáveis vão ter o mesmo valor e endereço
