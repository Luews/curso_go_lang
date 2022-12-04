package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!!!!!\n", "GO")

}

// comando pra ver a pagina do GO, offline: "godoc -http=:(sua porta)"
// um comando bem legal pra ver o que tem de errado no seu código é o: go vet (nome do arquivo)
// que no caso vai ser o ex: go vet comandos.go, tiver faltando alguma coisa ele vai te auxiliar e te mostrar
// um comando para ver aonde estão alocadas as informações gerais e instaladas do seu GO: go env
// o comando "go help" + algum comando, ex: go help get, vai te explicar o que o get faz
