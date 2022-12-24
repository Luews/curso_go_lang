// resumindo: existe um conceito, chamado moc
// você isola uma função, para fins de teste mesmo
// por exemplo, digamos que uma função x aciona uma função y, e a y faz uma query num banco de dados, mo role
// se eu isolar o X para não conversar com o y
// eu só quero ver a porra do x, isolando ele

package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular o que você já sabe... :)
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
