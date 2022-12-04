// O exerício você fez, mas depois de quebrar bem a cabeça, com a ajuda, segue feito e os comentários abaixo para revisão

package main

import "fmt"

func notaParaConceito(n float64) string {

	switch {

	case n >= 9 && n <= 10:
		return "A" // O return significa que ele vai sair do método e RETORNAR esse valor
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"

	}

}
func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}

/* Sua 2º tentativa

package main

import "fmt"

func notaParaConceito(nota_para_média float64) string {

	var nota_para_média = int(n)

	switch {
	case nota_para_média >= 6:
		fmt.Println("O Aluno passou na média, e foi aprovado!!")
	case nota_para_média < 5:
		fmt.Println("O aluno infelizmente não alcançou a média e foi reprovado! :( ")
	default:
		return "Nota inválida"

	}

}

func main() {

	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(3))
}

 seu erro foi tentar fazer do 0, era o que você tentou fazer de íncio mesemo e provavelmente você ajustou errado, mas só pega o código de novo vai substituindo colocando case
 colocando o : e é isso, você foi certo na primeira tentativa, não desista!!, segue o código feito por você com ajuda */
