// maps aninhados(um dentro do outro) são maps perfeitos para organização
// o primeiro map vai ter uma chave e o valor vai ser outro map, que vai ter uma chave e um valor

package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": { // percebe-se que no 2º map, como é um map mesmo, você vai abrir outra chave, se atende e lembre-se disso
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		}, // fechado a chave e colocando a vírgula para acrescentar mais valores
		"J": {
			"José João": 11325.45, // "José joão = string, 11325.45 = float64"
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcsPorLetra, "P")                // aqui é bem descritivo, mas vai apagar a letra P e seus valores
	for letra, funcs := range funcsPorLetra { // para cada uma das letras, e funções, no alcance da função que e gente criou
		fmt.Println(letra, funcs)
	}
}
