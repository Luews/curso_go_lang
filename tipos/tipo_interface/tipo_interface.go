package main

import "fmt"

type curso struct { // por aqui nada demais
	nome string
}

func main() {
	var coisa interface{} // percebe-se que criamos uma variavel chamada coisa que é uma interface vazia, vamos usar ela de temmplate, e podemos rodar vários dados conforme o código
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{} // outra interface vazia

	var coisa2 dinamico = "Opa" // aqui vamos atribuir
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"} // aqui é importante, pegamos a variável coisa 2 e atribuimos ela com a nossa interface VAZIA, como ela ta vazia, não tem nada, não retorna nada
	fmt.Println(coisa2)                                        // vai atribuir essa string e alocar dentro e printa-la normalmante
}
