package main

import "testing"

func testAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if esperado != soma {
		t.Errorf("esperado '%d', resultado '%d'", soma, esperado)
	}
}
