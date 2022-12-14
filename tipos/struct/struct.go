// um struct, seria um agrupador de dados, uma estrutura

package main

import "fmt"

type produto struct { // percebe-se sem separar por vírgula
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())
	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
}

// é só olhar que você consegue entender
// é tipo uma classe, é igualzinho,
// você agrupa uns dados
