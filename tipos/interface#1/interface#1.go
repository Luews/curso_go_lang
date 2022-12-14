/* o que seria uma interface?

é um conceito simples, uma interface é protocolo ele só define as funções para que alguma STRUCT TEM que ter para ser usada e passada como parâmetro
enfim */

package main

import "fmt"

type imprimivel interface {
	toString() string // para sabermos se uma estrutura é imprimível ou não, ela precisa ter esse método associado a ela
}
type pessoa struct {
	nome      string
	sobrenome string
}
type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}
func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"} // a palavra coisa vai alocar uma struct dentro dela e ser usada como parametro
	fmt.Println(coisa.toString())
	imprimir(coisa) // percebe-se que a palavra imprimir virou uma função, no python é comum ter funções iguais a essa, mas aqui devemos criar uma interface antes expecificar detalhadamente para podermos no final só colocar a palavra(parametro) que também é uma interface
	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)
	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)
}

// "Antes nós tinhamos funções que trabalhavam com dados, agora podemos ter dados que trabalham com funções"
