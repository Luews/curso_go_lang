package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // aqui vai ser um and, trabalho 1 e trabalho2
	comprarTv32 := trab1 != trab2    // ou exclusivo
	comprarSorvete := trab1 || trab2 // Aqui vai ser o or padrão que você já conhece
	return comprarTv50, comprarTv32, comprarSorvete
}
func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete) /* aqui nesse caso explicando, teve 2 parâmetros, que os 2 foram verdadeiros, ele conseguiu o trabalho1 e o trabalho2
	, sendoa assim, ele conseguiu comprar a Tv de 50 polegadas, e tomou sorvete, mas nisso não ficou tão saudável



	se mudarmos os parâmetros a resposta também, vai retornar outros valores booleanos de true ou false



	*/

}
