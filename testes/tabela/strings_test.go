package strings

import "testing" // sempre importar essa biblioteca, pra sinalizar

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)." //sempre fazer uma mensagem de erro, saber onde você errou, isso facilita muito, antes isso que você mesmo criou do que uma porra gigante que você não vai ler e se tentar não vai conseguir saber onde foi

func TestIndex(t *testing.T) { // sempre usar esse ponteiro aqui 
	t.Parallel()
	testes := []struct {
	texto string
	parte string
	esperado int
	}{
	{"Cod3r é show", "Cod3r", 0},
	{"", "", 0},
	{"Opa", "opa", -1},
	{"leonardo", "o", 2},
	}
	for _, teste := range testes {
	t.Logf("Massa: %v", teste)
	atual := strings.Index(teste.texto, teste.parte)
	if atual != teste.esperado {
	t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
	}
	}
	}
	69