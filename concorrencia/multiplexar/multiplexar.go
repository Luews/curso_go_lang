/* aqui não importa, mas a ideia dessa aula é pegar um canal Público, de outro código, e usar ele
dentro de outro canal, bem interessante, era pra funcionar esse tipo de import, mas o curso
ta desatulizado, talvez um dia eu fiquei esperto o suficiente pra atualizar essa porra e deixar
funcionardo */

// mas no geral e resumindo bem porcamente, é literalmente um canal dentro do outro, e esse canal interno
// é de outro código, outro arquivo, outra pasta, como se fosse uma biblioteca mesmo

package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {

	c := make(chan string) // canal c
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c

}
func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
