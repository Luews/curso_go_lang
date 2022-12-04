// é interessante porque você pode selecionar os tipos de dados que vão passar, bem interessante

package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string { // essa função eu tenho a capacidade de receber tipos de métodos, com vários tipos de dados sem causar um erro de compilação
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}
func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now())) // aqui era uma função que o software não estava esperando, por isso deu não sei
}
