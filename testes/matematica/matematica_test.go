package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

func TestMedia(t *testing.T) {
	valorExperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorExperado {
		t.Errorf(erroPadrao, valorExperado, valor)
	}
}
