// nesse test serve para ver se a maquina suporta esse testes

// a vida do programador é 20% programando e o resto testando e arrumando

// seria bom pegar um curso mais atualizado nessa porra, e rever esses conceitos de novo

// mas vendo por cima tendo uma noção já ta bom

package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
