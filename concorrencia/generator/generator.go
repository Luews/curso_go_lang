package main

// essa função generator, ele cria gourotines prontas e canais prontos pra você
// sem se preocupar em ficar administrando os canais, e mexendo e criando gourotines,

//pelo o que eu entendi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns
// <-chan - canal somente-leitura

func titulo(urls ...string) <-chan string { // a sintaxe é assim de um generator | Cria o que você quiser e mete um channel de leitura retornando

	c := make(chan string)
	for _, url := range urls { // aqui ignorando o indice
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1] // percebe-se que você não fica trabalhando com aquelas porras de canal, tem alguns já prontos aleatórios para serem utilizados
		}(url)
	}
	return c
}
func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
