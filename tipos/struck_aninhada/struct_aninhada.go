package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}
type pedido struct {
	userID int
	itens  []item // percebe-se que nessa struct tem outra struct em forma de slice
}

func (p pedido) valorTotal() float64 { // sempre usar essa variável p
	total := 0.0
	for _, item := range p.itens { // aqui um array
		total += item.preco * float64(item.qtde)
	}
	return total
}
func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 12.10}, // o vscode notificou que não precisa colocar o item na frente dos itens de slice, só abror e fechar as chaves e separar com vírgula ta bom
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal()) // printf para formatação

}
