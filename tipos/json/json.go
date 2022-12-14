// Se o json é trabalhoso e teria que ficar convertendo e desconvertendo, por que o uso dele?
// O json é normalmente bastante útil usando em webservices
// no back-end, no mongo db, é importante saber

package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"` // para json usamos aspas normal e depois aspas duplas
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1) // lembrando que esse json.Marshal é um método que retorna 2 valores, o primeiro valor, seria o próprio json e o 2º seria o erro, usando o _ para não executar o erro que não vamos precisar dele
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString :=
		`{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2) // se json.Marshal te retorna em struct, o json.Unmarshal pega o json para formatar em struct
	fmt.Println(p2.Tags[1])                 // aqui a gente selecionou umas das tags, mas poderiamos pegar a string inteira
}
