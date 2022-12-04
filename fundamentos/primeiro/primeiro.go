// Programas executáveis inciam pelo pacote main, caso esse pacote não exista, o programa não será executado
package main

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/

import "fmt"

// A porta de entrada de um programa Go é a função main
func main() {
	fmt.Print("Primeiro")
	fmt.Print("Programa")
	fmt.Print("Hello World")
}

// Resumindo antes: não faça comentários desnecessários, apenas que agregem sem estorvar

/*

1) Priorize código legível e faça comentários que agrega valor!
2) Evite comentários óbvios
3) Durante o curso abuse dos comentários

óbvio que durante o curso você não deve se restringir, porém lembrando que não deve adotar o costume de ANORTAR TUDO, uma situação real, na empresa seria ruim */
