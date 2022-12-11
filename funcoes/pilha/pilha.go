// Pelo o que eu entendi essa função de pilha serve para grandes projetos onde várias funções se chamam dentro das outras
// esse debug print stack, serve para printar o processo de execução dessas funções
/* é como se fosse uma pilha mesmo,

F3>
  F2>
    F1>
	 MAIN

Fazendo a pilha dessa forma, nós vemos que uma vai precisar executar a outra, e nesse debugprintstack, vemos o trajeto de como funciona se executarmos

*/
package main

import "runtime/debug"

func f3() {
	debug.PrintStack()
}
func f2() {
	f3()
}
func f1() {
	f2()
}
func main() {
	f1()
}
