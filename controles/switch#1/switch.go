// O Switch serve para lidar com seleções múltiplas

//.

//você vai usar o Switch ao inves de fazer os ifs e else ifs
// vai ficar mais "limpo" o código fazendo dessa forma e fica mais objetivo também,o switch foi feito pra lidar com esse tipo de problema

package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // serve para continuar executando os debaixo até concluir esse bloco
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: // o default serve quando nenhum dos cases é encontrado
		return "Nota inválida"
	}
}
func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))

}
