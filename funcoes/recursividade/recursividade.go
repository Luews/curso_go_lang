package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil // nil de nulo, pra não executar o erro
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}
func main() {
	resultado, _ := fatorial(5) // o _ para ignorar o erro, GO é meio chatinho com isso, precisa colocar os 2 parâmetros presentes
	fmt.Println(resultado)
	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	// como provalvemente você vai receber esses dados do usuário, você colocaria isso tudo num bloco só, trataria o resultado, e o erro, mesma sintaxes

	// Uma solução melhor seria... uint!
}
