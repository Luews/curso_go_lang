// o init, consegue rodar sem estar dentro da função main

// extremamente útil quando tem vários arquivos na mesma pasta, com isso, você não pode ter vários mains, mas pode ter vários inits, para diferenciar

package main

import "fmt"

func init() {
	fmt.Println("Inicializando2...")
}
