/* Dentro da estrutura If em Go é possível você ter um BLOCO de inicialização, antes de você executar essa
estrutura e esse bloco de inicialização vai te dar uma VARIÁVEL local dentro do escopo do BLOCO If */

/* Resumindo você vai poder criar uma variável dentro do if do nada, antes de declarar e inicializar ela */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // biblioteca de pegar a hora atual, e o nano segundo
	r := rand.New(s)                           // rand, aleatório
	return r.Intn(10)                          // colocar um range até 10
}
func main() {
	if i := numeroAleatorio(); i > 5 { // percebe-se o ponto e vírgula para separar as senteças // tb suportado no switch
		fmt.Println("Ganhou!!!") // caso o nanosegundo seja antes dos 10
	} else {
		fmt.Println("Perdeu!") // caso contrário
	}
}
