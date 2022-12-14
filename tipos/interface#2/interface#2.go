package main

import "fmt"

type esportivo interface {
	ligarTurbo() // percebe-se que a função ligar turbo não tem nenhum parâmetro e não retorna nada, ou seja, precisamos ALTERA-LA DE FORA....
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

//precisamos de um método onde juntamos, associamos, a struct ferrari com a interface esportivo

func (f *ferrari) ligarTurbo() { // aqui criamos um reciever junto com um ponteiro, porque vou precisar mexer na variável ligar turbo, se eu não tivesse um ponteiro e acionasse a ferraria, eu iria pegar uma cópia e com a cópia não ia conseguir alterar
	f.turboLigado = true // como o método ligarTurbo não tem nada, de fato agora a ferrari respeita a interface esportivo por conta que você colocou esses limites
}

func main() {

	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo() // percebe-se que eu utilizando a variável carro1 e acionando a função ligar turbo, ele vai modificar uma variável dentro da struct

	var carro2 esportivo = &ferrari{"F41", false, 0} //quando você tem um uso assim, a partir de uma INTERFACE, e você precisa passar a refêrencia da estrutura utilizando o "&"
	carro2.ligarTurbo()                              // vai mudar o turboLigado de false, para true com essa função da interface

	fmt.Println(carro1, carro2)

}
