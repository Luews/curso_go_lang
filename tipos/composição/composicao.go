// é muito fácil você ter uma composição de interface, você ter uma interface que é o resultado de uma COMPOSIÇÃO de OUTRAS INTERFACES
// é muito simples

package main

import "fmt"

type esportivo interface { // ele vai ter uma função ligarTurbo, mas de novo sem parâmetro e sem retorno, como se fosse um template para ser usado
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoluxuoso interface { // que é a junção das duas, e com isso carrega os códigos, etc delas
	esportivo
	luxuoso

	// pode adicionar novos métodos aqui dentro do escopo
}

type bwm7 struct{} // não vai ter nada aqui, mas a gente vai colocar usando o receiver

func (b bwm7) ligarTurbo() {
	fmt.Println("Ligou o turbo......")
}
func (b bwm7) fazerBaliza() {
	fmt.Println("Manobrou... estacionou!") // de forma implítica, ela respeita as 3 interfaces, a esportivo, luxuoso, e esportivoluxuoso
}

func main() {

	var b esportivoluxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
	// aqui da pra ver CLARAMENTE, que ele usou a interface que junta as duas, e usou funções distintas, de interfaces difernete, porém estavam juntas nesse interface main

}
