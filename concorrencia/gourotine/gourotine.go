package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)                                 // vai executar dando pausa de segundo em segundo
		fmt.Printf("%s: %s(iteração %d)\n", pessoa, texto, i+1) // usar o printf de formatação
	}

}

// o que é thread?
// thread é uma parte do código que pode ser executada de forma independentemente do programa principal
// outra linha de execução para ser de forma independete e com o intuito de ser rapida

func main() {

	/*
			fale("Doutor bélico", "Doutor bélico", 1)
			fale("Doutor bélico", "praaaaa", 1)
			fale("Doutor bélico", "Doutor bélico", 1)
			fale("Doutor bélico", "Quem?!", 1)
			fale("Doutor bélico", "bélico", 2)
			fale("Doutor bélico", "Quem que é?", 1)
			fale("Doutor bélico", "EU SOU O DOUTOR BÉLICO", 1)
			fale("Doutor bélico", "maxa", 2)
			fale("Doutor bélico", "Aí piranha, você já viu um doutor?!", 1)

			fmt.Println("----------------------------------------------------------------------")

			fale("Lucas", "é a PORRA do Dr Bélico, ta?", 5)

		// vai ser executado de forma padrão, segundo por segundo
		// é um processo serial, ele vai executar na leitura padrão, e depois o outro

		go fale("Doutor bélico", "Aí piranha, você já viu um doutor?", 500)
		go fale("Manoel Gomes", "aiiiwnnn", 500)

		time.Sleep(time.Second * 5)

		fmt.Println("Fim") */

	// uma gourotine, só vai continuar executando, se a thread principal estiver funcionando

	go fale("Doutor bélico", "Aí piranha você já viu um doutor?!", 10)
	go fale("Doctor bélico", "bélico", 4)
	fale("Dr bélico", "Só opioide receitado", 5)

	// vai ser uma salada, mas vai sair de forma independete, os gourotinees, ou as threads, devem ser colocadas primeiro e sempre pra fechar, uma thread da linha main deve ser acionada no final

}
