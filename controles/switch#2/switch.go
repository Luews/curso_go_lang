/* Se você abrir um switch, e não tiver nenhuma condição associada a ele, ele vai procurar algum CASE
que seja VERDADEIRO, o primeiro case que ele achar e for verdadeiro, ele vai executar

muito bom seloco */

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite.")
	}
}
