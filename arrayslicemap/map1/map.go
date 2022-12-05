/* para que serve o map no golang? : Um map é uma estrutura de dados capaz de ASSOCIAR uma CHAVE a um VALOR.
Eles são extremamente eficientes em retornar um valor baseado na chave,
isso porque os maps são uma implementação de uma hash table */

package main

import "fmt"

func main() {
	// var aprovados map[int]string , nós não podemos criar um mapa assim
	// mapas devem ser inicializados
	aprovados := make(map[int]string) // deve ser feito assim, criar e inicializar
	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // fpr cpf (int) , nome (string) no alcance de aprovados
		fmt.Printf("%s (CPF: %d)\n", nome, cpf) // na mesma forma também, \n para ir quebrando a linha, usando o for para ir executando cada um de uma vez, e passando o nome e o cpf, na ordem que você quer printar e como ponteiro/
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682) // para deletar alguma informação, usar o método delete assim, nas próximas linhas ele não vai puxar/bucar/printar
	fmt.Println(aprovados[95135745682])
}
